package controllers

type AdminController struct {
	BaseController
}

func (this *AdminController) Edit() {
	this.TplName = "edit.html"
}
