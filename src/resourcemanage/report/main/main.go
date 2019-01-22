package main

import (
	"fmt"
	"github.com/astaxie/beego"
	_ "go_dev/day13/beego_example/router"
	"os"
)

func init() {
	fmt.Println("a")
}
func main() {

	environ := os.Environ()
	fmt.Println(environ)
	beego.Run()
}
