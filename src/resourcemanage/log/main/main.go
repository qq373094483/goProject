package main

import (
	"fmt"
	"github.com/astaxie/beego"
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
