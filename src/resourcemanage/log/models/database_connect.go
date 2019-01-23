package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql")

func init(){
	orm.Debug=true
	orm.RegisterDriver("mysql",orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:qq123456@tcp(127.0.0.1:3306)/log?charset=utf8",1000,2000)
	orm.RegisterModel(new(CommonLog),new(ExceptionLog))
	orm.RunSyncdb("default", false, true)
}