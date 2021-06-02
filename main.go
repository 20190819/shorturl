package main

import (
	"shorturl/controllers"
	_ "shorturl/routers"

	"github.com/astaxie/beego"
)

func main() {
	beego.Router("/", &controllers.IndexController{}, "get:Index")
	beego.Run()
}
