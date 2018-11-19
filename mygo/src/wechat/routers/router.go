package routers

import (
	"wechat/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.ArticleController{}, "get:IndexList")
	beego.Router("/subject/?:id", &controllers.ArticleController{}, "get:SubjectList")
	beego.Router("/class", &controllers.ArticleClassController{}, "get:Get")
	beego.Router("/comment/:id", &controllers.CommentController{})
}
