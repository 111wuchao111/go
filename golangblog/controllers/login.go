package controllers

import (
	"golangblog/models"

	"github.com/astaxie/beego/cache"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/utils/captcha"
)

type LoginController struct {
	BaseController
}

var cpt *captcha.Captcha

func init() {
	// use beego cache system store the captcha data
	store := cache.NewMemoryCache()
	cpt = captcha.NewWithFilter("/captcha/", store)
}

func (this *LoginController) Login() {
	var user models.User
	if this.GetSession("user") != nil {
		this.Redirect("/admin", 302)
	}
	uname := this.Input().Get("username")
	upwd := this.Input().Get("pwd")
	if uname == "" || upwd == "" || cpt.VerifyReq(this.Ctx.Request) == false {
		this.TplName = "login.html"
		return
	}
	o := orm.NewOrm()
	o.QueryTable("user").Filter("username", uname).One(&user)
	if user.Password != models.GetMd5(upwd) {
		this.TplName = "login.html"
		return
	}
	this.SetSession("user", user.Username)
	this.Redirect("/admin", 302)
}
