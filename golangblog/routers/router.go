package routers

import (
	"golangblog/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/index", &controllers.MainController{})
	beego.Router("/time", &controllers.MainController{}, "get:Time")
	beego.Router("/about", &controllers.MainController{}, "get:About")
	beego.Router("/detail/:id", &controllers.MainController{}, "get:Detail")

	beego.Router("/controller", &controllers.UeditorController{}, "*:ControllerUE")
	beego.Router("/admin/login", &controllers.LoginController{}, "*:Login")
	beego.Router("/admin", &controllers.AdminController{}, "*:Index")
	beego.Router("/admin/list", &controllers.AdminController{}, "*:List")
}
