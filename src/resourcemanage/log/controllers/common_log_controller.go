package controllers

import (
	"encoding/json"
	"fmt"
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
	var ob models.CommonLog
	cc:=p.Ctx.Input.RequestBody
	json.Unmarshal(cc,&ob)
	a:=p.GetString("Id")
	fmt.Println(a)
	m["code"]=200
	m["message"]="success"
	p.Data["json"]=m
	p.ServeJSON()
}