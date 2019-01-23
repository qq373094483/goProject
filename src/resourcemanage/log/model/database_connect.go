package model

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql")

func init(){
	orm.Debug=true
	orm.RegisterDriver("mysql",orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:qq123456@tcp(127.0.0.1:3306)/log?charset=utf8")
	orm.SetMaxIdleConns("default",1000)
	orm.SetMaxOpenConns("default",2000)
}