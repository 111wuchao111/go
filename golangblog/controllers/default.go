package controllers

import (
	"golangblog/models"

	"strconv"

	"github.com/astaxie/beego/orm"
)

//默认controller
type MainController struct {
	BaseController
}

//首页
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

//时间轴
func (c *MainController) Time() {
	var articles []models.Article
	o := orm.NewOrm()
	o.QueryTable("article").OrderBy("-Time").Limit(10).All(&articles)
	c.Data["Articles"] = articles
	c.TplName = "time.html"
}

//关于我
func (c *MainController) About() {
	var user models.User
	orm.NewOrm().QueryTable("user").Filter("id", 1).One(&user)
	c.Data["User"] = user
	c.TplName = "about.html"
}

func (c *MainController) Detail() {
	var article models.Article
	var watch_desc_article []models.Article
	id := c.Ctx.Input.Param(":id")
	intid, _ := strconv.Atoi(id)
	o := orm.NewOrm()
	o.QueryTable("article").RelatedSel().Filter("id", intid).One(&article)
	o.QueryTable("article").RelatedSel().OrderBy("-Watch_count").Limit(10).All(&watch_desc_article)
	c.Data["Watch_desc_article"] = watch_desc_article
	c.Data["Article"] = article
	c.TplName = "detail.html"
}
