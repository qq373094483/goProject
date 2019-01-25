package main

import ("github.com/astaxie/beego"
	"resourcemanage/log/common/format"
	_ "resourcemanage/log/routers"
)

func main() {
	format.A()
	beego.Run()
}
