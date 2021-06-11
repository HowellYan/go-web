package main

import (
	"github.com/beego/beego/v2/server/web"
	_ "webapp/routers"
)

func main() {
	web.Run()
}
