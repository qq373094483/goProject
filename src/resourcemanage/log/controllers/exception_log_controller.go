package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"resourcemanage/log/constant"
	"resourcemanage/log/models"
)

type  ExceptionLogController struct{
	beego.Controller
}


func (this *ExceptionLogController) Insert() error{
	logs.Debug("enter index controller")
	exceptionLog :=new(models.ExceptionLog)

	if err:=json.Unmarshal(this.Ctx.Input.RequestBody, exceptionLog);err!=nil{
		return err
	}
	ormer:=orm.NewOrm()
	exceptionLog.ExceptionLogInsert(&ormer,"defalut")
	this.Data["json"]=constant.Messages[constant.Success]
	this.ServeJSON()
	return nil
}