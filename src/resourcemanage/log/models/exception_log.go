package models

import (
	"encoding/json"
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

type exceptionLogAlias ExceptionLog
type exceptionLog struct {
	ErrorInfo string `json:"errorInfo"`
	GmtCreate string `json:"gmtCreate"`
	*exceptionLogAlias
}
func (this *ExceptionLog) UnmarshalJSON(data []byte) error {
	aux := exceptionLog{
		exceptionLogAlias: (*exceptionLogAlias)(this),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	date,err:=time.ParseInLocation("2006-01-02 15:04:05",aux.GmtCreate,time.Local)
	if err!=nil {
		return err
	}
	this.GmtCreate=date
	this.ErrorInfo=aux.ErrorInfo
	return nil
}

func (this *ExceptionLog) MarshalJSON() ([]byte, error) {
	return json.Marshal(exceptionLog{
		exceptionLogAlias:  (*exceptionLogAlias)(this),
		GmtCreate: this.GmtCreate.Format("2006-01-02 15:04:05"),
		ErrorInfo:this.ErrorInfo,
	})
}