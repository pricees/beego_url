package controllers

import (
	"crypto/md5"
	"encoding/base64"
	"github.com/astaxie/beego"
)

type ShortenController struct {
	beego.Controller
}

func shortenMe(str string) (short string) {
	hasher := md5.New()
	hasher.Write([]byte(str))
	short = base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	return
}

func (this *ShortenController) Post() {
	this.Layout = "layout.html"
	longurl := this.GetString("longurl")
	this.Data["longurl"] = longurl
	this.Data["shorturl"] = shortenMe(longurl)
	this.TplNames = "shorten.get.tpl"
}

func (this *ShortenController) Get() {
	this.Layout = "layout.html"
	this.TplNames = "shorten.get.tpl"
}
