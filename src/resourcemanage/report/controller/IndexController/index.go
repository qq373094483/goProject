package IndexController

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type IndexController struct {
	beego.Controller
}

func (p *IndexController) Index() {
	logs.Debug("enter index controller")
	p.TplName = "IndexController/index.html"
	m := make(map[string]interface{})
	m["code"] = 200
	m["message"] = "success"
	p.Data["json"] = m
	//p.ServeJSON()

	/*p.Data["xml"]=m
	p.ServeXML()*/
}
