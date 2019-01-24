package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"resourcemanage/log/models"
)

type  CommonLogController struct{
	beego.Controller
}



func (p *CommonLogController) Insert(){
	logs.Debug("enter index controller")
	m:=make(map[string]interface{})
	commonLog:=new(models.CommonLog)
	json.Unmarshal(p.Ctx.Input.RequestBody, commonLog)
	m["code"]=200
	m["message"]="success"
	p.Data["json"]=m
	p.ServeJSON()
}