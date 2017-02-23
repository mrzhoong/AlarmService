package g

import (
	"github.com/astaxie/beego"
	"strconv"
	"sync"
)

type GlobalConfig struct {
	Debug        bool
	RedisAddr    string
	RedisMaxIdle int
	DbHost       string
	DbPort       string
	DbName       string
	DbUser       string
	DbPwd        string
}

var (
	// 配置文件名称
	config     *GlobalConfig
	configLock = new(sync.RWMutex)
)

func Config() *GlobalConfig {
	configLock.RLock()
	defer configLock.RUnlock()
	return config
}

func LoadConfig() {
	var c GlobalConfig
	MaxIdle := beego.AppConfig.String("redismaxidle")
	c.RedisAddr = beego.AppConfig.String("redisaddr")
	c.RedisMaxIdle, _ = strconv.Atoi(MaxIdle)
	c.DbHost = beego.AppConfig.String("dbhost")
	c.DbPort = beego.AppConfig.String("dbport")
	c.DbUser = beego.AppConfig.String("dbuser")
	c.DbPwd = beego.AppConfig.String("dbpwd")
	c.DbName = beego.AppConfig.String("dbname")
	config = &c
}
