package controllers

import (
	"github.com/astaxie/beego"
)

type IndexController struct {
	beego.Controller
}

func (_this *IndexController) Index() {
	_this.Ctx.Output.Body([]byte("shortUrl-demo"))
}
