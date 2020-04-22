package models

import (
	_ "github.com/GoAdminGroup/go-admin/modules/db/drivers/mysql"
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

	//注册model
	orm.RegisterModel(new(Users), new(UserProfile), new(Regions))

	//var w io.Writer
	//orm.DebugLog = orm.NewLog(w)

	_ = orm.RunSyncdb(dbConnection, false, true) //同步数据表


}
