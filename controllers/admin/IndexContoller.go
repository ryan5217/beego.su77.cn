package admin

type IndexController struct {
	BaseController
}

func (this *IndexController) Get() {

	this.TplName = "/admin/index.html"
}
