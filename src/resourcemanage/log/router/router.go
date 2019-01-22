package router

import (
	"github.com/astaxie/beego"
	"go_dev/day13/beego_example/controller/IndexController"
)

func init() {
	beego.Router("/index", &IndexController.IndexController{}, "*:Index")
}
