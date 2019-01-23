package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type ExceptionLog struct {
	Log
	ErrorInfo string `db:"error_info"`
}

func (exceptionLog *ExceptionLog) ExceptionLogInsert(ormer *orm.Ormer,databaseAliasName string){
	(*ormer).Using(databaseAliasName)
	exceptionLog.GmtCreate=time.Now()
	exceptionLog.Id=0
	(*ormer).Insert(exceptionLog)
}