package main

import (
	_ "beego.su77.cn/routers"
	_ "github.com/GoAdminGroup/go-admin/adapter/beego"
	_ "github.com/GoAdminGroup/go-admin/modules/db/drivers/mysql"
	_ "github.com/GoAdminGroup/themes/adminlte"
	"github.com/astaxie/beego/plugins/cors"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import your used driver
	"time"
)

func init() {
	//数据库链接
	dbDriver := beego.AppConfig.String("DB_DRIVER")
	dbConnection := beego.AppConfig.String("DB_CONNECTION")
	dbHost := beego.AppConfig.String("DB_HOST")
	dbPort := beego.AppConfig.String("DB_PORT")
	dbDatabase := beego.AppConfig.String("DB_DATABASE")
	dbUsername := beego.AppConfig.String("DB_USERNAME")
	dbPassword := beego.AppConfig.String("DB_PASSWORD")
	dbPrefix := beego.AppConfig.String("DB_PREFIX")

	dsn := dbUsername +":"+ dbPassword +"@tcp("+ dbHost +":"+ dbPort +")/"+ dbDatabase

	_ = orm.RegisterDataBase(dbConnection, dbDriver, dsn)
	orm.DefaultTimeLoc = time.UTC
	orm.RegisterModelWithPrefix(dbPrefix)
	orm.Debug = true

	//var w io.Writer
	//orm.DebugLog = orm.NewLog(w)

	_ = orm.RunSyncdb(dbConnection, false, true) //同步数据表
}

func main() {
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		//允许访问所有源
		AllowAllOrigins: true,
		//可选参数"GET", "POST", "PUT", "DELETE", "OPTIONS" (*为所有)
		//其中Options跨域复杂请求预检
		AllowMethods:   []string{"*"},
		//指的是允许的Header的种类
		AllowHeaders: 	[]string{"*"},
		//公开的HTTP标头列表
		ExposeHeaders:  []string{"X-Requested-With, Content-Type, X-Token-Auth, Authorization, Content-Length"},
		//如果设置，则允许共享身份验证凭据，例如cookie
		AllowCredentials: true,
	}))

	beego.Run()

	//goAdmin
	//app := beego.NewApp()
	//
	//eng := engine.Default()
	//
	//dbHost := beego.AppConfig.String("DB_HOST")
	//dbPort := beego.AppConfig.String("DB_PORT")
	//dbDatabase := beego.AppConfig.String("DB_DATABASE")
	//dbUsername := beego.AppConfig.String("DB_USERNAME")
	//dbPassword := beego.AppConfig.String("DB_PASSWORD")
	//
	//cfg := config.Config{
	//	Databases: config.DatabaseList{
	//		"default": {
	//			Host:       dbHost,
	//			Port:       dbPort,
	//			User:       dbUsername,
	//			Pwd:        dbPassword,
	//			Name:       dbDatabase,
	//			MaxIdleCon: 50,
	//			MaxOpenCon: 150,
	//			Driver:     config.DriverMysql,
	//		},
	//	},
	//	Store: config.Store{
	//		Path:   "./uploads",
	//		Prefix: "uploads",
	//	},
	//	UrlPrefix:   "go_admin",
	//	IndexUrl:    "/",
	//	Debug:       true,
	//	Language:    language.CN,
	//	ColorScheme: adminlte.ColorschemeSkinBlack,
	//}
	//
	//adminPlugin := admin.NewAdmin(datamodel.Generators).AddDisplayFilterXssJsFilter()
	//
	//template.AddComp(chartjs.NewChart())
	//
	//// add generator, first parameter is the url prefix of table when visit.
	//// example:
	////
	//// "user" => http://localhost:9087/admin/info/user
	////
	//adminPlugin.AddGenerator("user", datamodel.GetUserTable)
	//adminPlugin.AddGenerator("users", goAdmin.GetUsersTable)
	//
	//// customize a plugin
	//
	//examplePlugin := example.NewExample()
	//
	//// load from golang.Plugin
	////
	//// examplePlugin := plugins.LoadFromPlugin("../datamodel/example.so")
	//
	//// customize the login page
	//// example: https://github.com/GoAdminGroup/go-admin/blob/master/demo/main.go#L30
	////
	//// template.AddComp("login", datamodel.LoginPage)
	//
	//// load config from json file
	////
	//// eng.AddConfigFromJSON("../datamodel/config.json")
	//
	//beego.SetStaticPath("/uploads", "uploads")
	//
	//if err := eng.AddConfig(cfg).AddPlugins(adminPlugin, examplePlugin).Use(app); err != nil {
	//	panic(err)
	//}
	//
	//// you can custom your pages like:
	//
	//app.Handlers.Get("/admin", func(ctx *context.Context) {
	//	eng.Content(ctx, func(ctx interface{}) (types.Panel, error) {
	//		return datamodel.GetContent()
	//	})
	//})
	//
	////beego.BConfig.Listen.HTTPAddr = "127.0.0.1"
	////beego.BConfig.Listen.HTTPPort = 9087
	//go app.Run()
	//
	//quit := make(chan os.Signal)
	//signal.Notify(quit, os.Interrupt)
	//<-quit
	//log.Print("closing database connection")
	//eng.MysqlConnection().Close()
}

