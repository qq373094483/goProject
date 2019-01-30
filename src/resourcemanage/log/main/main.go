package main

import (
	"github.com/astaxie/beego"
	_ "resourcemanage/log/routers"
)

func main() {
	/*u := models.CommonLog{models.Log{GmtCreate: time.Now()}}
	uj, _ := json.Marshal(u)
	uu:= models.CommonLog{}
	json.Unmarshal(uj, &uu)
	fmt.Println(u.GmtCreate)
	fmt.Println(uu.GmtCreate)*/
	beego.Run()
}
