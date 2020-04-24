package middlewares

import (
	"beego.su77.cn/controllers/admin"
	"errors"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func AdminAuthHandler() {

	beego.InsertFilter("/admin/*",beego.BeforeRouter, func(context *context.Context) {

		if context.Request.RequestURI != "/admin/login" {

			if context.Request.RequestURI != "/admin/register" {

				//此处可以校验一下ip，设备等
				//todo 需要修改 优化代码
				token := context.Request.Header.Get("authorization")

				defer func() {
					//捕获到异常错误 返回信息401未授权
					if r := recover(); r != nil {

						context.Redirect(302, "/admin/login")

					}
				}()

				if len(admin.CheckToken(token)) == 0 {
					context.Redirect(302, "/admin/login")
					panic(errors.New("user stop run"))
				}
			}
		}
	})

}
