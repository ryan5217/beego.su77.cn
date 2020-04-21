package goAdmin

type IndexController struct {
	BaseController
}

func (this *IndexController) Get() {
	this.TplName = "portal/index.html"
}
