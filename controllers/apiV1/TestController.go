package apiV1

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

type DataController struct {
	beego.Controller
}

type GetGetController struct {
	beego.Controller
}

type LIKE struct {
	Food string `json:"food"`
	Watch string `json:"watch"`
	Listen string `json:"listen"`
}

type JSONS struct {
	//必须的大写开头
	Code string `json:"code"`
	Message  string `json:"message"`
	User []string `json:"user"`//key重命名,最外面是反引号
	Like LIKE `json:"like"`
}

//Get()就是默认函数
func (c *DataController) Get() {
	data := &JSONS{"100", "获取成功",
		[]string{"maple","18"},LIKE{"蛋糕","电影","音乐"}}
	c.Data["json"] = data
	c.ServeJSON()
}

func (c *DataController) Test() {
	c.TplName = "apiV1/test.tpl"
}

func (c *DataController) List() {
	data := &JSONS{"100", "获取成功",
		[]string{"maple","19"},LIKE{"蛋糕","电影","音乐"}}
	c.Data["json"] = data
	c.ServeJSON()
}


func (c *MainController) Get() {
	//c.Data["Website"] = "beego.me"
	//c.Data["Email"] = "astaxie@gmail.com"
	c.Data["Content"] = "value123678"
	c.TplName = "apiV1/test.tpl"
	//c.ServeJSON()
	//fmt.Println("test")
}



