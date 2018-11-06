package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type User struct {
	Id       int
	Name     string`orm:"unique"`//用户名唯一
	Password string
	Article []*Article `orm:"rel(m2m)"` //user和article的关系是多对多
} //用户结构体
type Article struct {
	Id             int                                           //id
	Articletitle   string    `orm:"size(20)"`                    //标题
	Articletime    time.Time `orm:"auto_now_add;type(datetime)"` //时间
	Articleimg     string                                        //路径
	Articlecount   int `default(0);null`                         //阅读次数,初始为0，允许为空
	Articlecontent string                                        //文章内容
	Articletype    *Articletype `orm:"rel(fk)"`                  //设置外键，一对多的关系
	User         []*User      `orm:"reverse(many)"`
} //文章结构体
type Articletype struct {
	Id       int
	Typename string      `orm:"size(20)"`
	Article [] *Article `orm:"reverse(many)"` //相对应的,多对多
} //类型表

func init() {
	orm.RegisterDataBase("default", "mysql", "root:@163@tcp(127.0.0.1:3306)/myproject?charset=utf8&loc=Asia%2FShanghai")
	orm.RegisterModel(new(User), new(Article), new(Articletype)) //创建表
	orm.RunSyncdb("default", false, true)                        //别名，是否强制更新，是否终端可见
}
