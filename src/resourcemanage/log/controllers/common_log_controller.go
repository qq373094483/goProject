package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"resourcemanage/log/constant"
	"resourcemanage/log/models"
)

type  CommonLogController struct{
	beego.Controller
}


func (this *CommonLogController) Insert(){
	logs.Debug("enter index controller")
	commonLog:=new(models.CommonLog)
	json.Unmarshal(this.Ctx.Input.RequestBody, commonLog)
	ormer:=orm.NewOrm()
	commonLog.CommonLogInsert(&ormer,"defalut")
	this.Data["json"]=constant.Messages[constant.Success]
	this.ServeJSON()
}