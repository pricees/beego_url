package controllers

import (
	"crypto/md5"
	"encoding/base64"
	"github.com/astaxie/beego"
)

var (
  shortyLookup *beego.BeeCache
)

const DONT_EXPIRE = 0

func init() {
  shortyLookup = beego.NewBeeCache()
  shortyLookup.Every = DONT_EXPIRE
  shortyLookup.Start()
}

type ShortenController struct {
	beego.Controller
}

func storeMe(shorturl string, longurl string) {
  if !shortyLookup.IsExist(shorturl) {
    shortyLookup.Put(shorturl, longurl, DONT_EXPIRE)
    shortyLookup.Put(longurl, shorturl, DONT_EXPIRE)
  }
}

func shortenMe(str string) (short string) {
	hasher := md5.New()
	hasher.Write([]byte(str))
	short = base64.URLEncoding.EncodeToString(hasher.Sum(nil))[0:5]
	return
}

func (this *ShortenController) Post() {
	this.Layout = "layout.html"
	longurl := this.GetString("longurl")
  
  var shorturl string
  if shortyLookup.IsExist(longurl) {
    shorturl = shortyLookup.Get(longurl).(string)
  } else {
    shorturl = shortenMe(longurl)
    storeMe(shorturl, longurl)
  }
	this.Data["longurl"] = longurl
	this.Data["shorturl"] = "http://localhost:8080/" + shorturl
	this.TplNames = "shorten.get.tpl"
}

func (this *ShortenController) Get() {
	this.Layout = "layout.html"
	this.TplNames = "shorten.get.tpl"
}

/*
  Redirect Handler
*/
type RedirectController struct {
  beego.Controller
}

func (this *RedirectController) Get() {
  shorturl := this.Ctx.Params[":shorturl"]

  if shortyLookup.IsExist(shorturl) {
    longy := shortyLookup.Get(shorturl).(string)
    this.Redirect(longy, 302)
  } else {
    this.Redirect("/", 302)
  }
}
