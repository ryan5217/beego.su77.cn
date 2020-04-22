package routers

import (
	"beego.su77.cn/controllers"
	"beego.su77.cn/controllers/admin"
	"beego.su77.cn/controllers/apiV1"
	"beego.su77.cn/controllers/goAdmin"
	"beego.su77.cn/controllers/portal"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func init() {
    beego.Router("/", &controllers.MainController{})

    //beego.Router()
    beego.Router("/test", &apiV1.MainController{})
    beego.Router("/data", &apiV1.DataController{})
	beego.Router("/data/h", &apiV1.DataController{})
    beego.Router("/data/test", &apiV1.DataController{}, "get:Test")
    beego.Router("/data/list", &apiV1.DataController{}, "get:List")

    beego.Any("/any", func(context *context.Context) {
		var beego beego.Controller

    	data := make(map[string] interface{})
    	data["nihao"] = "nihao"
    	data["nibuhao"] = "nibuhao"

		beego.Data["json"] = data
		beego.ServeJSON()


		//context.Output.Body([]byte("notAllowed"))
    	//context.Output.Body([]byte("asdasd"))
	})

    //admin登录test
    ns_admin :=
    	beego.NewNamespace("/admin",
			beego.NSRouter("/login", &admin.LoginController{}, "get:Login;post:DoLogin"),
			beego.NSRouter("/register", &admin.LoginController{}, "get:Register;post:DoRegister"),
			beego.NSRouter("/index", &admin.IndexController{}),
			beego.NSRouter("/curl", &admin.LoginController{}, "get:Curl"),
			beego.NSRouter("/region", &admin.RegionController{}, "get:Index"),
			beego.NSRouter("/values", &admin.RegionController{}, "get:Values"),
			beego.NSRouter("/curd", &admin.RegionController{}, "get:Curd"),
			beego.NSRouter("/select", &admin.RegionController{}, "get:Select"),
			beego.NSRouter("/get_all_regions", &admin.RegionController{}, "get:GetAllRegions"),
		)

	ns_goadmin :=
		beego.NewNamespace("/go_admins",
			beego.NSRouter("/index", &goAdmin.IndexController{}),
			beego.NSAny("/any", func(context *context.Context) {
				var beegoInt beego.Controller


				data := make(map[string] interface{})
				data["nihao"] = "nihao"
				data["nibuhao"] = "nibuhao"

				beegoInt.Data["json"] = data
				beegoInt.ServeJSON()


				//context.Output.Body([]byte("notAllowed"))
				//context.Output.Body([]byte("asdasd"))
			}),
		)

	ns_home :=
		beego.NewNamespace("/portal",
			beego.NSRouter("/index", &portal.IndexController{}, "get:Get"),
		)
	//注册 namespace
	beego.AddNamespace(ns_admin, ns_goadmin, ns_home)


}
