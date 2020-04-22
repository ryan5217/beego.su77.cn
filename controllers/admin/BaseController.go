package admin

import "github.com/astaxie/beego"

const TimeDefault string = "2006-01-02 15:04:05"

type BaseController struct {
	beego.Controller
}

type Data struct {
	Code int `json:"code"`
	Message string `json:"message"`
	Data interface{} `json:"data"`
}
