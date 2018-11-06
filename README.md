beego_project
====
![](https://img.shields.io/cocoapods/l/Alamofire.svg?style=flat)
[![](https://travis-ci.org/Alamofire/Alamofire.svg?branch=master)](https://travis-ci.org/Alamofire/Alamofire)

![](https://img.shields.io/badge/-v1.0.0-519dd9.svg)


概览:
基于beego 的后台管理系统


性能卓越

MVC

简单好用

参见

`beego 官方网站：`
https://beego.me



-----
文件目录
```
conf
    --app.conf//配置文件
controllers
    --default.go//业务逻辑处理
models
    --model.go //数据库业务逻辑
routers
    --router.go//路由控制器
static 
    --.......//静态文件等相关
test
    --default_test.go//测试文件
views   
    --......//视图文件
main.go //主程序入口
```
依赖需求
----
```
1. beego   


go get github.com/astaxie/beego

2. ORM

go get github.com/astaxie/beego/orm

3. go-sql-driver

go get github.com/go-sql-driver/mysql

4. convey

go get github.com/smartystreets/goconvey/convey

```

使用方法
----
`cd go PATH`

`go get github.com/wangzhengzhi/beego_project`

`cd beego_project`

` bee run`

注意事项：
----

*1默认使用mysql，需要自己创建数据库

如下文件可以配置
```
models
    ---model.go
```
如下图代码
```
func init() {
	orm.RegisterDataBase("default", "mysql", "root:password@tcp(127.0.0.1:3306)/database_name?charset=utf8&loc=Asia%2FShanghai")
	orm.RegisterModel(new(User), new(Article), new(Articletype)) //创建表
	orm.RunSyncdb("default", false, true)                        //别名，是否强制更新，是否终端可见
}



```
`password`                  数据库密码

`database_name`             数据库 需要手动创建

`&loc=Asia%2FShanghai`              默认使用ASIA ShangHai 时区，如果不设置会出现时区错误问题，请注意




2 配置文件


```

conf
    ---app.conf

```
可以更改运行模式以及运行端口


后续支持
====
1 将会支持微信支付/支付宝支付等接口

2 支持redis 


3 更改前端代码，更加美观









