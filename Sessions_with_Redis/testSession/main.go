package main

import (
	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/session/redis"
	_ "testSession/routers"
)

func main() {
	beego.Run()
}
