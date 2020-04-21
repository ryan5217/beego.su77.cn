package portal

import "github.com/astaxie/beego"

type IndexController struct {
	beego.Controller
}

func (c *IndexController) Get() {

	c.Data["title"] = "你好 beego2020"

	c.TplName = "portal/index.html"
}