package routers

import (
	"github.com/astaxie/beego"
	"myproject/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/register", &controllers.MainController{}) //register路由
	//有了自定义的get请求，不会再去访问默认的post
	beego.Router("/login", &controllers.MainController{}, "get:ShowLogin;post:HandleLogin")
	beego.Router("/index", &controllers.MainController{}, "get:ShowIndex;post:HandleIndex")
	beego.Router("/addArticle", &controllers.MainController{}, "get:ShowArticle;post:HandleArticle") //post处理下拉框
	beego.Router("/content", &controllers.MainController{}, "get:ShowContent;post:HandleContent")
	beego.Router("/update", &controllers.MainController{}, "get:ShowUpdate;post:HandleUpdate")
	beego.Router("/delete", &controllers.MainController{}, "get:HandleDelete")                                         //删除的路由
	beego.Router("/addArticleType", &controllers.MainController{}, "get:ShowAddArticleType;post:HandleAddArticleType") //添加分类的路由

}
