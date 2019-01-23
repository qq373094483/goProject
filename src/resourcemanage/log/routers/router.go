package routers

import (
	"github.com/astaxie/beego"
	"resourcemanage/log/controllers"
)

func init() {
	beego.Router("/index", &controllers.CommonLogController{}, "post:Insert")
}
