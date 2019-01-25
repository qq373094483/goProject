package json

import "time"

type JsonDate time.Time
type JsonTime time.Time

func  UnmarshalJSON(data []byte) (time.Time,error) {
	local, err := time.ParseInLocation(`"2006-01-02"`, string(data), time.Local)
	return time.Time(local),err
}

func (p *JsonTime) UnmarshalJSON(data []byte) error {
	local, err := time.ParseInLocation(`"2006-01-02 15:04:05"`, string(data), time.Local)
	*p = JsonTime(local)
	return err
}
