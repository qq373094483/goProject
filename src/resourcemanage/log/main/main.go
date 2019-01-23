package main

import (
	"github.com/astaxie/beego/orm"
	"resourcemanage/log/model"
)

func main() {
	o:=orm.NewOrm()
	o.Using("default")
	common:=model.CommonLog{Log:model.Log{}}
	common.Insert()
}
