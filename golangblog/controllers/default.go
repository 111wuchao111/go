package controllers

import (
	"golangblog/models"

	"github.com/astaxie/beego/orm"
)

type MainController struct {
	BaseController
}

func (c *MainController) Get() {
	var brands []models.Brand
	var kerns []models.Article
	var articles []models.Article
	var watch_desc_article []models.Article
	var user models.User
	o := orm.NewOrm()
	o.QueryTable("brand").Limit(3).All(&brands)
	o.QueryTable("article").Filter("Is_kern", 1).RelatedSel().Limit(2).All(&kerns)
	o.QueryTable("article").RelatedSel().OrderBy("-Time").Limit(10).All(&articles)
	o.QueryTable("article").RelatedSel().OrderBy("-Watch_count").Limit(10).All(&watch_desc_article)
	o.QueryTable("user").Filter("Id", 1).One(&user)
	c.Data["User"] = user
	c.Data["Watch_desc_article"] = watch_desc_article
	c.Data["Article_kerns"] = kerns
	c.Data["Articles"] = articles
	c.Data["Brand"] = brands
	c.TplName = "index.html"
}
