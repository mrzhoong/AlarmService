package models

// import (
// 	"github.com/Unknwon/com"
// 	"github.com/astaxie/beego/orm"
// 	"github.com/mattn/go-sqlite3"
// )

// const (
// 	// 设置数据库路径
// 	_DB_NAME = "data/beeblog.db"
// 	// 设置数据库名称
// 	_SQLITE3_DRIVER = "sqlite3"
// )

//  策略
type Strategy struct {
	Strategy_id         int64
	Strategy_name       string
	Strategy_conditions string
	Sending_method      int64
	Send_to             int64
}

// 告警源

// 用户
type User struct {
	Id                int64
	Name              string
	Tel               string
	Email             string
	Wechat            string
	Organization_code int64
	Organization_name string
}

// func RegisterDB() {
// 	// 检查数据库文件
// 	if !com.IsExist(_DB_NAME) {
// 		os.MkdirAll(path.Dir(_DB_NAME), os.ModePerm)
// 		os.Create(_DB_NAME)
// 	}

// 	// 注册模型
// 	orm.RegisterModel(new(Category), new(Topic), new(Comment))
// 	// 注册驱动（“sqlite3” 属于默认注册，此处代码可省略）
// 	orm.RegisterDriver(_SQLITE3_DRIVER, orm.DR_Sqlite)
// 	// 注册默认数据库
// 	orm.RegisterDataBase("default", _SQLITE3_DRIVER, _DB_NAME, 10)
// }

// 告警条件
