package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "wuchao.com"
	c.Data["Email"] = "314237418.com"
	c.TplName = "index.html"
}
