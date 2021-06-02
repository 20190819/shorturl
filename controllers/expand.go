package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
)

type ExpandController struct {
	beego.Controller
}

func init() {
	urlcache, _ = cache.NewCache("memory", `{"interval:0"}`)
}

func (_this *ExpandController) ToLong() {
	var result ShortResult
	shorturl := _this.Input().Get("shorturl")
	result.Data.UrlShort = shorturl

	if urlcache.IsExist(shorturl) {
		if value, ok := urlcache.Get(shorturl).(string); ok {
			result.Data.UrlLong = value
		}
	}
	_this.Data["json"] = result
	_this.ServeJSON()
}
