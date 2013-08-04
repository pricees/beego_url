package main

import (
	"beego_url/controllers"
	"github.com/astaxie/beego"
)

func main() {
	beego.SetStaticPath("/images", "static/images")
	beego.SetStaticPath("/css", "static/css")
	beego.SetStaticPath("/js", "static/js")

	beego.Router("/:shorturl:string", &controllers.RedirectController{})
	beego.Router("/", &controllers.ShortenController{})
	beego.Run()
}
