package controllers

import (
	"github.com/astaxie/beego"
)

type ArticleController struct {
	beego.Controller
}

//获取首页文章列表
func (this *ArticleController) IndexList() {
	this.Data["json"] = map[string]interface{}{"success": 0, "message": "111"}
	this.ServeJSON()
	return
}

//专题列表
func (this *ArticleController) SubjectList() {
	this.Data["json"] = map[string]interface{}{"success": 0, "message": "22222"}
	this.ServeJSON()
	return
}
