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
	if this.Ctx.Input.Method() == "GET" {
		var categorys []models.Category
		o := orm.NewOrm()
		o.QueryTable("article_category").All(&categorys)
		this.Data["Categorys"] = categorys
		this.TplName = "admin_add_article.html"
		return
	}
	this.TplName = "admin_add_article.html"

}

//文章详情

//保存文章

//图片上传
func (this *AdminController) Upload() {
	f, h, _ := this.GetFile("view_img_url")
	path := "/home/wuchao/mygo/src/golangblog/static/upload/" + h.Filename
	f.Close()
	this.SaveToFile("view_img_url", path)
	this.Data["json"] = map[string]interface{}{"state": "SUCCESS", "url": "/static/upload/" + h.Filename}
	this.ServeJSON()
}
