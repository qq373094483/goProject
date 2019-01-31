package models

import (
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