package routers

import (
	"shorturl/controllers"

	"github.com/astaxie/beego"
)

func init() {
	nsApi := beego.NewNamespace("/api",
		beego.NSNamespace("/shorten",
			beego.NSRouter("/", &controllers.ShortController{}, "get:ToShort"),
		),
		beego.NSNamespace("/expand",
			beego.NSRouter("/", &controllers.ExpandController{}, "get:ToLong"),
		),
	)

	// 注册 namespace
	beego.AddNamespace(nsApi)
}
