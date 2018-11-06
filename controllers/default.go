package controllers

import (
	"math"
	"myproject/models"
	"path"
	"strconv"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	//orm的插入
	/*o := orm.NewOrm()
	user := models.User{}
	user.Name = "hello"
	user.Password = "1234"
	_, err := o.Insert(&user)
	if err != nil {
		beego.Info(err)
		return
	}*/
	//orm的查询
	/*o := orm.NewOrm()
	user := models.User{}*/
	/*user.Id = 1

	err := o.Read(&user)
	if err != nil {
		beego.Info(err)
		return
	}*/
	/*user.Name="hello"
	err:=o.Read(&user,"Name")
	if err!=nil {
		beego.Info(err)
		return

	}
	beego.Info("查询成功")
	beego.Info(user.Name)*/
	//orm的更新
	/*o := orm.NewOrm()
	user := models.User{}
	user.Id = 1
	err := o.Read(&user)
	if err == nil {
		user.Name = "helloworld"
		user.Password = "111"
		_,err:=o.Update(&user)
		if err!=nil {
			beego.Info(err)
			return

		}
	}*/
	/*//orm的删除
	o:=orm.NewOrm()
	user:=models.User{}
	user.Id=1
	_,err:=o.Delete(&user)
	if err!=nil {
		beego.Info(err)
		return

	}

	c.Data["data"] = "今天中午吃饺子"
	c.TplName = "test.html"*/
	c.TplName = "register.html"
}
func (c *MainController) Post() {

	username := c.GetString("username")
	password := c.GetString("password")
	beego.Info(username, password)
	if username == "" || password == "" {
		beego.Info("数据不能为空")
		c.Redirect("/register", 302)
		return
	}
	o := orm.NewOrm()
	user := models.User{}
	user.Name = username
	user.Password = password

	_, err := o.Insert(&user)
	if err != nil {
		beego.Info("插入数据库失败")
		c.Redirect("register.html", 302)
		return
	}
	c.Redirect("login.html", 302) //跳转到登陆界面
}
func (c *MainController) ShowLogin() {
	c.TplName = "login.html"
}
func (c *MainController) HandleLogin() {
	/*c.Ctx.WriteString("登陆成功，post请求")*/
	username := c.GetString("username")
	password := c.GetString("password")
	//beego.Info(username, password)
	if username == "" || password == "" {
		beego.Info("数据不合法")
		c.TplName = "login.html"
		return
	}
	o := orm.NewOrm()
	user := models.User{}
	user.Name = username

	err := o.Read(&user, "Name") //服务器查询用户名是否失败
	if err != nil {
		beego.Info("没有此用户名")
		c.TplName = "login.html"
		return
	}
	if user.Name == username && user.Password == password { //密码校验
		c.Redirect("/index", 302) //登陆成功之后跳转到index用户页面
	} else {
		c.Ctx.WriteString("登陆失败,请重新登陆")

	}

}
func (c *MainController) ShowIndex() {

	pageindex, err := strconv.Atoi(c.GetString("pageindex"))
	//c.getstring是array，从HTML中拿取数据,默认值存在不可能失败，存在error
	if err != nil {
		pageindex = 1 //如果没有获取内容那么则是返回默认值1
	}
	o := orm.NewOrm()
	var article []models.Article
	qs := o.QueryTable("Article") //类似select * from Article
	count, err := qs.Count()      // 返回条目数量
	if err != nil {
		beego.Info("查询条目失败", err)
		return
	}
	pagesize := 1 //每一页显示的条数数量
	sumpage := float64(count) / float64(pagesize)
	c.Data["count"] = count                //一共有多少条目数量
	c.Data["sumpage"] = math.Ceil(sumpage) //一共有多少页

	start := pagesize * (pageindex - 1)
	qs.Limit(pagesize, start).All(&article) //限制每一页显示的数量使用limit()
	c.Data["article"] = article
	c.Data["pageindex"] = pageindex
	//beego.Info("pageindex", pageindex)
	//首页和末页的逻辑判断，让首页不能有上一页，末页不能有下一页
	//处理上一页和下一页出现的时机，不能在首页或者末尾出现的东西
	firstpage := false
	if pageindex == 1 { //标识是否是首页
		firstpage = true
	}
	c.Data["firstpage"] = firstpage
	lastpage := false
	if pageindex == int(math.Ceil(sumpage)) { //标识是否是末页
		lastpage = true
	}
	c.Data["lastpage"] = lastpage
	//获取类型数据
	var articletype []models.Articletype
	o.QueryTable("Articletype").All(&articletype)
	c.Data["articletype"] = articletype

	c.TplName = "index.html"
}
func (c *MainController) HandleIndex() {
	//处理下拉框
	//接收数据
	selectname := c.GetString("select") //拿到html中的数值
	beego.Info("selectname:", selectname)
	//处理数据
	if selectname == "" {
		beego.Info("下拉框传递数据失败")
		return
	}
	//查询数据
	o := orm.NewOrm()
	var article []models.Article
	o.QueryTable("Article").RelatedSel("Articletype").Filter("Articletype__Typename", selectname).All(&article)
	beego.Info(article)

}
func (c *MainController) ShowArticle() {

	c.TplName = "add.html"
}
func (c *MainController) HandleArticle() {
	//拿到数据
	articletitle := c.GetString("articletitle")
	content := c.GetString("content")
	//beego.Info("articletitle:", articletitle)
	//beego.Info("content:", content)

	f, h, err := c.GetFile("uploadname")
	defer f.Close()
	//限定格式
	fileext := path.Ext(h.Filename) //拿到文件后缀名字
	//beego.Info("文件后缀是：", fileext)
	if fileext != ".jpg" && fileext != ".png" {
		beego.Info("上传文件类型错误")
		return
	}
	//限定大小
	if h.Size > 50000000 { //h.size是以bit为单位的
		beego.Info("上传文件过大")
		return
	}
	//处理文件名字
	filename := time.Now().Format("2006.01.02.15.04.05") + fileext //当前时间加文件后缀
	if err != nil {
		beego.Info("上传失败", err)
		return

	} else {
		c.SaveToFile("uploadname", "./static/img/"+filename)
	}
	//判断合法性
	if articletitle == "" || content == "" {
		beego.Info("更新文章数据失败")
		return

	}
	// 插入数据
	o := orm.NewOrm()
	article := models.Article{}
	article.Articlecontent = content
	article.Articletitle = articletitle
	article.Articleimg = "/static/img/" + filename //地址，可以不要点
	_, err = o.Insert(&article)                    //插入
	if err != nil {
		beego.Info(err, "插入数据库错误")
	}
	//返回
	c.Redirect("/index", 302)
}
func (c *MainController) ShowContent() {
	id, err := c.GetInt("id")
	beego.Info("id is", id)
	if err != nil {
		beego.Info("获取id失败", err)
		return
	}
	o := orm.NewOrm()
	article := models.Article{Id: id}
	err = o.Read(&article)
	if err != nil {
		beego.Info("查询article失败", err)
		return
	}
	c.Data["article"] = article
	c.TplName = "content.html"

}
func (c *MainController) HandleContent() {

}
func (c *MainController) ShowUpdate() {
	id, err := c.GetInt("id")
	//beego.Info("id is", id)
	if err != nil {
		beego.Info("获取id失败", err)
		return
	}
	o := orm.NewOrm()
	article := models.Article{Id: id}
	err = o.Read(&article)
	if err != nil {
		beego.Info("查询article失败", err)
		return
	}
	c.Data["article"] = article
	c.TplName = "update.html"
}
func (c *MainController) HandleUpdate() {
	id, _ := c.GetInt("id")
	articletitle := c.GetString("articletitle")
	content := c.GetString("content")
	f, h, err := c.GetFile("uploadname")
	defer f.Close()
	fileext := path.Ext(h.Filename) //拿到后缀
	if fileext != ".jpg" && fileext != ".png" {
		beego.Info("文件类型错误")
		return
	}
	if h.Size > 50000000 {
		beego.Info("文件过大")
		return
	}
	//给文件重新命名
	filename := time.Now().Format("2006.01.02.15.04.05") + fileext
	if err != nil {
		beego.Info("上传失败", err)
		return

	} else {
		c.SaveToFile("uploadname", "./static/img/"+filename)
	}
	//判断数据合法性
	if articletitle == "" || content == "" {
		beego.Info("添加文章数据错误")
		return
	}
	// 更新操作
	o := orm.NewOrm()
	article := models.Article{Id: id}
	err = o.Read(&article)
	if err != nil {
		beego.Info("查询数据错误")
		return

	}
	article.Articlecontent = content
	article.Articletitle = articletitle
	article.Articleimg = "/static/img/" + filename                              //地址，可以不要点
	_, err = o.Update(&article, "Articlecontent", "Articletitle", "Articleimg") //选择性更新数据
	if err != nil {
		beego.Info(err, "插入数据库错误")
		return
	}
	c.Redirect("/index", 302)

}
func (c *MainController) HandleDelete() {
	//拿到数据
	id, err := c.GetInt("id")
	if err != nil {
		beego.Info("Id数据获取失败")
		return
	}
	//执行删除操作
	o := orm.NewOrm()
	article := models.Article{Id: id}
	err = o.Read(&article)
	if err != nil {
		beego.Info("查询数据失败")
		return
	}
	o.Delete(&article) //删除数据

	//返回列表页面
	c.Redirect("/index", 302)
}
func (c *MainController) ShowAddArticleType() {
	o := orm.NewOrm()
	var articletype []models.Articletype //文章类型是一个集合
	//查询
	_, err := o.QueryTable("Articletype").All(&articletype)
	if err != nil {
		beego.Info("查询类型错误")
		//暂时不return

	}
	c.Data["articletype"] = articletype
	c.TplName = "addType.html"
}
func (c *MainController) HandleAddArticleType() { //处理添加类型业务
	// 获取数据
	typeName := c.GetString("typeName")
	if typeName == "" {
		beego.Info("添加类型为空")
		return

	} //判断数据是否合法
	o := orm.NewOrm()
	var articletype models.Articletype
	articletype.Typename = typeName
	_, err := o.Insert(&articletype)
	if err != nil {
		beego.Info("插入失败")
		return

	} //执行插入操作
	c.Redirect("/addArticleType", 302)
	//返回视图

}
