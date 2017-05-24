package g

import (
	"encoding/json"
	"strconv"
	"sync"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/toolkits/file"
)

type GlobalConfig struct {
	Debug        bool
	ConfPath     string
	MailTopic    string
	RedisAddr    string
	RedisMaxIdle int
	DbHost       string
	DbPort       string
	DbName       string
	DbUser       string
	DbPwd        string
}

type StrategyList struct {
	Id   int64  `json:id`
	Name string `json:name`
}

type SendMethod struct {
	Id   int64  `json:id`
	Name string `json:name`
}

type SendTo struct {
	Id     int64  `json:id`
	Name   string `json:name`
	Mobile string `json:mobile`
	WeChat string `json:wechat`
	Email  string `json:email`
}

type StrategyInfo struct {
	//	策略ID
	Id int64 `json:id`
	//	发送的组信息
	EquipGroup []StrategyList `json:strategylist`
	//	策略名称
	Name string `json:name`
	//	条件，暂未处理
	Condition string `json:condition`
	//	发送方式
	Method []SendMethod `json:sendmethod`
	//	发送人员
	To []SendTo `json:sendto`
}

type Strategy struct {
	Count     int64          `json:count`
	Strategys []StrategyInfo `json:strategys`
}

var (
	// 配置文件名称
	config     *GlobalConfig
	configLock = new(sync.RWMutex)
	// cfg json
	strategys    *Strategy
	strategyLock = new(sync.RWMutex)
)

func Config() *GlobalConfig {
	configLock.RLock()
	defer configLock.RUnlock()
	return config
}
func Cfg() *Strategy {
	strategyLock.RLock()
	defer strategyLock.RUnlock()
	return strategys
}

func LoadConfig() {
	var c GlobalConfig
	MaxIdle := beego.AppConfig.String("redismaxidle")

	c.Debug, _ = beego.AppConfig.Bool("debug")
	c.MailTopic = beego.AppConfig.String("mailtopic")
	c.RedisAddr = beego.AppConfig.String("redisaddr")
	c.RedisMaxIdle, _ = strconv.Atoi(MaxIdle)
	c.DbHost = beego.AppConfig.String("dbhost")
	c.DbPort = beego.AppConfig.String("dbport")
	c.DbUser = beego.AppConfig.String("dbuser")
	c.DbPwd = beego.AppConfig.String("dbpwd")
	c.DbName = beego.AppConfig.String("dbname")
	c.ConfPath = beego.AppConfig.String("path")

	config = &c
}

func LoadCfg(cfg string) {
	log := logs.GetLogger("LoadCfg")
	if cfg == "" {
		log.Println("use -c to specify configuration file")
	}

	if !file.IsExist(cfg) {
		log.Fatalln("config file:", cfg, "is not existent")
	}

	configContent, err := file.ToTrimString(cfg)
	if err != nil {
		log.Fatalln("read config file:", cfg, "fail:", err)
	}

	// var c Strategys
	var c Strategy
	err = json.Unmarshal([]byte(configContent), &c)
	if err != nil {
		log.Fatalln("parse config file:", cfg, "fail:", err)
	}
	strategyLock.Lock()
	defer strategyLock.Unlock()
	strategys = &c
	log.Println("read config file:", cfg, "successfully")
}

func SetLog() {
	logs.SetLogger(logs.AdapterFile, `{"filename":"alarm_service.log",
		"level":7,"maxlines":1000000,"maxsize":0,"daily":true,"maxdays":20}`)
	//	行号、文件名
	logs.EnableFuncCallDepth(true)
	logs.Async()
}
