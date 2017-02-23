package main

import (
	"AlarmService/controllers"
	"AlarmService/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	models.RegisterDB()
}

func main() {

	orm.Debug = true
	orm.RunSyncdb("default", false, true)
	beego.Router("/", &controllers.HomeController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/strategy", &controllers.StrategyController{})
	beego.AutoRouter(&controllers.StrategyController{})
	//	User
	beego.Router("/user", &controllers.UserController{})
	beego.AutoRouter(&controllers.UserController{})
	beego.Router("/eventlog", &controllers.LoginController{})
	beego.Router("/operationlog", &controllers.LoginController{})
	beego.Router("/config", &controllers.LoginController{})
	beego.Router("/test", &controllers.TestController{})
	beego.Router("/event", &controllers.EventController{})
	beego.Run()
}
