package controllers

import (
	"github.com/astaxie/beego"
)

type ErrorController struct {
	beego.Controller
}

type Data struct {
	Code int `json:"code"`
	Message string `json:"message"`
	Data interface{} `json:"data"`
}

func (this *ErrorController) ApiJsonReturn(code int, message string, data interface{})  {
	var json Data
	json.Code = code
	json.Message = message
	json.Data = data

	this.Data["json"] = json
	this.ServeJSON()
	this.StopRun()
}

func (c *ErrorController) Error401() {

	content := "未经授权，请求要求验证身份"
	data := make(map[string] interface{})
	c.ApiJsonReturn(401, content, data)
}

func (c *ErrorController) Error403() {

	content := "服务器拒绝请求"
	data := make(map[string] interface{})
	c.ApiJsonReturn(403, content, data)
}

func (c *ErrorController) Error404() {

	content := "很抱歉您访问的地址或者方法不存在"
	data := make(map[string] interface{})
	c.ApiJsonReturn(404, content, data)
}

func (c *ErrorController) Error500() {

	content := "server error"
	data := make(map[string] interface{})
	c.ApiJsonReturn(500, content, data)
}

func (c *ErrorController) Error503() {

	content := "服务器目前无法使用（由于超载或停机维护）"
	data := make(map[string] interface{})
	c.ApiJsonReturn(503, content, data)
}

func (c *ErrorController) ErrorApi() {
	content := "service error"
	data := make(map[string] interface{})
	c.ApiJsonReturn(500, content, data)
}
