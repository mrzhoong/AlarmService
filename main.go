package main

import (
	"AlarmService/controllers"
	//	"AlarmService/models"
	"github.com/astaxie/beego"
)

func main() {

	beego.Router("/", &controllers.HomeController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/strategy", &controllers.AlarmStrategy{})
	beego.Router("/strategy/add", &controllers.AlarmStrategy{}, "post:Add")
	beego.Router("/strategy/delete", &controllers.AlarmStrategy{}, "post:Delete")
	beego.Router("/eventlog", &controllers.LoginController{})
	beego.Router("/operationlog", &controllers.LoginController{})
	beego.Router("/config", &controllers.LoginController{})
	beego.Run()
}
