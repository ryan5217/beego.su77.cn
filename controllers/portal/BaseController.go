package portal

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

func (obj Data) Success(data interface{}) Data {
	obj.Code = 0
	obj.Message = "success"
	obj.Data = data
	return obj
}

func (obj Data) Error(message string) Data {
	obj.Code = 1
	obj.Message = message
	obj.Data = make(map[string]string)
	return obj
}
