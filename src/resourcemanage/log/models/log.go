package models

import (
	"encoding/json"
	"time"
)

type Log struct {
	Id int64 `db:"id"`
	SystemIdentify string `orm:"column(system_identify)"`
	Param string `db:"param"`
	MethodSignature string `db:"method_signature"`
	Uuid string `db:"uuid"`
	GmtCreate time.Time
}


type logAlias Log
type commonLog struct {
	GmtCreate string `json:"gmtCreate"`
	*logAlias
}

func (log *Log) UnmarshalJSON(data []byte) error {
	aux := commonLog{
		logAlias: (*logAlias)(log),
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
		logAlias:  (*logAlias)(log),
		GmtCreate: log.GmtCreate.Format("2006-01-02 15:04:05"),
	})
}