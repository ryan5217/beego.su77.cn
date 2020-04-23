package middlewares

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"strings"
)

func AdminAuthHandler() {

	beego.InsertFilter("/*",beego.BeforeRouter, func(context *context.Context) {
		if context.Request.RequestURI != "/admin/login" {

			if context.Request.RequestURI != "/admin/register" {

				//此处可以校验一下ip，设备等
				//todo 需要修改 优化代码
				token := context.Request.Header.Get("authorization")

				fmt.Print(strings.Index(context.Request.RequestURI,"/admin/login"))

				if strings.Index(context.Request.RequestURI,"/admin/login") >= 0 || strings.Index(context.Request.RequestURI,"/static") >= 0 {
					//过滤不要控制的
				} else if token == "" {
					_, _ = context.ResponseWriter.Write([]byte("您无权访问"))
				}
			}
		}
	})

}
