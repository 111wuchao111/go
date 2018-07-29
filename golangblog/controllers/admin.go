package controllers

import (
	"golangblog/models"

	"github.com/astaxie/beego/orm"
)

type AdminController struct {
	BaseController
}

func (this *AdminController) Prepare() {
	userSession := this.GetSession("user")
	if userSession == nil {
		this.Redirect("/admin/login", 302)
	}
	this.Layout = "admin_layout.html"
}

//后台首页(文章列表)
func (this *AdminController) Index() {
	var article_list []models.Article
	o := orm.NewOrm()
	o.QueryTable("article").RelatedSel().OrderBy("-Time").Limit(10).All(&article_list)
	this.Data["Article_list"] = article_list
	this.TplName = "admin_list.html"
}

//文章列表
func (this *AdminController) List() {
	this.TplName = "admin_list.html"
}

//添加文章
func (this *AdminController) Add() {

}

//文章详情

//保存文章
