package routers

import (
	"github.com/astaxie/beego"
	"resourcemanage/log/controllers"
)

func init() {
	beego.Router("/commom_log", &controllers.CommonLogController{}, "post:Insert")
	beego.Router("/exception_log", &controllers.ExceptionLogController{}, "post:Insert")
}
