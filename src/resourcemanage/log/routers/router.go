package routers

import (
	"github.com/astaxie/beego"
	"resourcemanage/log/controllers"
)

func init() {
	beego.Router("/common_log", &controllers.CommonLogController{}, "post:Insert")
	beego.Router("/exception_log", &controllers.ExceptionLogController{}, "post:Insert")
}
