package model

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type CommonLog struct {
	Log
}
func init(){
	orm.RegisterModel(new(CommonLog))
}

func (commonLog *CommonLog) Insert() (o orm.Ormer){
	o=orm.NewOrm()
	o.Using("default")
	commonLog.GmtCreate=time.Now()
	commonLog.Id=0
	o.Insert(commonLog)
	return
}