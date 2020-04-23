package portal

type IndexController struct {
	BaseController
}

func (c *IndexController) Get() {

	c.Data["strong"] = "@ryan"
	c.Data["title"] = "你好 beego2020"

	c.TplName = "portal/index.html"
}