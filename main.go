package main

import (
	"github.com/astaxie/beego"
	_ "myproject/models"
	_ "myproject/routers"
)

func main() {
	beego.AddFuncMap("ShowPrePage", HandlePrePage)
	beego.AddFuncMap("ShowNextPage", HandleNextpage) //两个funmap需要在beego运行之前跑起来
	beego.Run()
}
func HandlePrePage(data int) int {
	pageindex := data - 1
	return pageindex

} //处理HTML中的前一页面的函数
func HandleNextpage(data int) int {
	pageindex := data + 1
	return pageindex
} //处理HTML中的后一页面的函数
