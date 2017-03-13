package main

import (
	"AlarmService/controllers"
	"AlarmService/g"
	"AlarmService/models"
	"AlarmService/redis"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	// "log"
	// "os"
	// "os/signal"
	// "syscall"
)

func init() {
	g.LoadConfig()
	g.LoadCfg("D:\\Development\\src\\AlarmService\\cfg.json")
	redis.InitConnPool()
	models.RegisterDB()

	// sigs := make(chan os.Signal, 1)
	// signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	// go func() {
	// 	<-sigs
	// 	redis.ConnPool.Close()
	// 	log.Println("get signal close:", sigs)
	// 	os.Exit(0)
	// }()

	// log.Println("get signal close:", sigs)
}

func main() {

	orm.Debug = false
	orm.RunSyncdb("default", false, true)

	beego.Router("/", &controllers.HomeController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/strategy", &controllers.StrategyController{})
	beego.AutoRouter(&controllers.StrategyController{})
	//	User
	beego.Router("/user", &controllers.UserController{})
	beego.AutoRouter(&controllers.UserController{})
	beego.Router("/eventlog", &controllers.RecordController{})
	beego.Router("/operationlog", &controllers.OperationController{})
	beego.Router("/config", &controllers.ConfigController{})
	beego.AutoRouter(&controllers.ConfigController{})
	beego.Router("/test", &controllers.TestController{})
	beego.Router("/event", &controllers.EventController{})
	beego.Run()

	select {}
}
