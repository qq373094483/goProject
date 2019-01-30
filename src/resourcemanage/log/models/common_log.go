package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type CommonLog struct {
	Log
}
func (commonLog *CommonLog) CommonLogInsert(ormer *orm.Ormer,databaseAliasName string){
	(*ormer).Using(databaseAliasName)
	commonLog.GmtCreate=time.Now()
	commonLog.Id=0
	(*ormer).Insert(commonLog)
	return
}

