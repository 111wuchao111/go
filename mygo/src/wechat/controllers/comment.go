package controllers

import (
	"github.com/astaxie/beego"
)

type CommentController struct {
	beego.Controller
}

//获取某个文章的评论
func (this *CommentController) Get() {
	this.Data["json"] = map[string]interface{}{"success": 0, "message": "wwww"}
	this.ServeJSON()
	return
}

//提交某个文章的评论
func (this *CommentController) Post() {

}
