package model

import "github.com/astaxie/beego/orm"


type ExceptionLog struct {
	Log
	ErrorInfo string `db:"error_info"`
}

func init(){
	orm.RegisterModel(new(ExceptionLog))
}