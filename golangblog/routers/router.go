package routers

import (
	"golangblog/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/time", &controllers.MainController{}, "get:Time")
	beego.Router("/about", &controllers.MainController{}, "get:About")
}
