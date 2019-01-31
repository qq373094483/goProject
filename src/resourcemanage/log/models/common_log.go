package models

import (
	"encoding/json"
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

type commonLogAlias Log
type commonLog struct {
	GmtCreate string `json:"gmtCreate"`
	*commonLogAlias
}

func (log *Log) UnmarshalJSON(data []byte) error {
	aux := commonLog{
		commonLogAlias: (*commonLogAlias)(log),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	date,err:=time.ParseInLocation("2006-01-02 15:04:05",aux.GmtCreate,time.Local)
	if err!=nil {
		return err
	}
	log.GmtCreate=date
	return nil
}

func (log *Log) MarshalJSON() ([]byte, error) {
	return json.Marshal(&commonLog{
		commonLogAlias: (*commonLogAlias)(log),
		GmtCreate:      log.GmtCreate.Format("2006-01-02 15:04:05"),
	})
}