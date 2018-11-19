package controllers

import (
	"fmt"
	"wechat/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type ArticleClassController struct {
	beego.Controller
}

//得到文章类别列表
func (this *ArticleClassController) Get() {
	var aricleClass []models.ArticleClass
	o := orm.NewOrm()
	o.QueryTable("article_class").All(&aricleClass)
	fmt.Println(aricleClass)
	this.Data["json"] = map[string]interface{}{"status": "success", "data": aricleClass}
	this.ServeJSON()
}
