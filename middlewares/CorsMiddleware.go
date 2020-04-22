package middlewares

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
)

func CorsHandler() {
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		//允许访问所有源
		AllowAllOrigins: true,
		//可选参数"GET", "POST", "PUT", "DELETE", "OPTIONS" (*为所有)
		//其中Options跨域复杂请求预检
		AllowMethods:   []string{"*"},
		//指的是允许的Header的种类
		AllowHeaders: 	[]string{"*"},
		//公开的HTTP标头列表
		ExposeHeaders:  []string{"X-Requested-With", "Content-Type", "X-Token-Auth", "Authorization", "Content-Length"},
		//ExposeHeaders:  []string{"X-Requested-With, Content-Type, X-Token-Auth, Authorization, Content-Length"},
		//如果设置，则允许共享身份验证凭据，例如cookie
		AllowCredentials: true,
	}))
}