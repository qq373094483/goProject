package format

import (
	"container/list"
	"fmt"
	"gopkg.in/fatih/set.v0"
	"time"
)

//使用Mysql日期格式化表示法
const (
	a  = "%a" //缩写星期名
	b  = "%b" //缩写月名
	c  = "%c" //月，数值
	dd = "%D" //带有英文前缀的月中的天
	d  = "%d" //月的天，数值(00-31)
	e  = "%e" //月的天，数值(0-31)
	f  = "%f" //微秒
	hh = "%H" //小时 (00-23)
	h  = "%h" //小时 (01-12)
	ii = "%I" //小时 (01-12)
	i  = "%i" //分钟，数值(00-59)
	j  = "%j" //年的天 (001-366)
	k  = "%k" //小时 (0-23)
	l  = "%l" //小时 (1-12)
	mm = "%M" //月名
	m  = "%m" //月，数值(00-12)
	p  = "%p" //AM 或 PM
	r  = "%r" //时间，12-小时（hh:mm:ss AM 或 PM）
	ss = "%S" //秒(00-59)
	s  = "%s" //秒(00-59)
	tt = "%T" //时间, 24-小时 (hh:mm:ss)
	uu = "%U" //周 (00-53) 星期日是一周的第一天
	u  = "%u" //周 (00-53) 星期一是一周的第一天
	vv = "%V" //周 (01-53) 星期日是一周的第一天，与 %X 使用
	v  = "%v" //周 (01-53) 星期一是一周的第一天，与 %x 使用
	ww = "%W" //星期名
	w  = "%w" //周的天 （0 = 星期日, 6 = 星期六）
	xx = "%X" //年，其中的星期日是周的第一天，4 位，与 %V 使用
	x  = "%x" //年，其中的星期一是周的第一天，4 位，与 %v 使用
	yy = "%Y" //年，4 位
	y  = "%y" //年，2 位
)

var formatSet = set.New(set.ThreadSafe)

func init() {
	formatSet.Add(a, b, c, dd, d, e, f, hh, h, ii, i, j, k, l, mm, m, p, r, ss, s, tt, uu, u, vv, v, ww, w, xx, x, yy, y)
}

func A() {
	fmt.Println(formatSet.Has("%Y"))
	fmt.Println(formatSet)
	err:=checkFormat("%Y %A")
	if err!=nil{
		fmt.Println(err)
	}
}

func Format(datestr, format string) (time.Time, error) {
	var s time.Time
	return s, nil
}

func Parse(time time.Time, format string) (string, error) {
	return "", nil
}

func checkFormat(format string) error {
	l := list.New()
	formatVal := ""
	for _, str := range format {
		fmt.Println(l.Len())
		if string(str) == "%" {
			l.PushBack(string(str))
			continue
		}
		if l.Len() == 1 {
			l.Remove(l.Back())
			formatVal = "%" + string(str)
		}
		if formatSet.Has(formatVal) {

		} else {
			return fmt.Errorf("data type '%s' not found", formatVal)
		}
	}
	return nil
}
