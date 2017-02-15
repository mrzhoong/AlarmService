package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
)

// const (
// 	// 设置数据库路径
// 	_DB_NAME = "data/beeblog.db"
// 	// 设置数据库名称
// 	_SQLITE3_DRIVER = "sqlite3"
// )

//  策略
type Strategy struct {
	Id int64
	// Strategy_id         int64
	Strategy_name       string
	Strategy_conditions string
	Sending_method      string
	Send_to             string
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

func RegisterDB() {
	dbhost := beego.AppConfig.String("dbhost")
	dbport := beego.AppConfig.String("dbport")
	dbuser := beego.AppConfig.String("dbuser")
	dbpwd := beego.AppConfig.String("dbpwd")
	dbname := beego.AppConfig.String("dbname")
	// 注册模型
	orm.RegisterModel(new(Strategy), new(User))
	// mysql 属于默认注册，此处代码可省略）
	orm.RegisterDriver("mysql", orm.DRMySQL)
	// 注册默认数据库
	conn := dbuser + ":" + dbpwd + "@tcp(" + dbhost + ":" + dbport + ")/" + dbname + "?charset=utf8"
	orm.RegisterDataBase("default", "mysql", conn)
	// orm.RegisterDataBase("default", "mysql", "root:123456@/default?charset=utf8")
}

// 获取策略列表
func GetAllStrategys() ([]*Strategy, error) {
	o := orm.NewOrm()
	strategys := make([]*Strategy, 0)

	qs := o.QueryTable("strategy")
	_, err := qs.All(&strategys)
	// strategys := []Strategy{Strategy{Strategy_id: 1, Strategy_name: "联通", Strategy_conditions: "实时", Sending_method: "微信", Send_to: "秦羽"},
	// 	Strategy{Strategy_id: 2, Strategy_name: "华为", Strategy_conditions: "实时", Sending_method: "短信", Send_to: "张三"},
	// 	Strategy{Strategy_id: 3, Strategy_name: "科技部", Strategy_conditions: "定时", Sending_method: "邮件", Send_to: "李四,王五"},
	// 	Strategy{Strategy_id: 4, Strategy_name: "中联值班组", Strategy_conditions: "实时", Sending_method: "电话", Send_to: "332"}}
	return strategys, err
}

// 增加策略
func AddStrategy(name string, methods string, conditions string, send_to string) error {
	// tidNum, err := strconv.ParseInt(id, 10, 64)
	// if err != nil {
	// 	beego.Error(err)
	// 	return err
	// }

	o := orm.NewOrm()
	strategy := &Strategy{
		Strategy_name:       name,
		Sending_method:      methods,
		Strategy_conditions: conditions,
		Send_to:             send_to,
	}
	_, err := o.Insert(strategy)
	if err != nil {
		return err
	}
	return nil
}

func GetStrategy(tid string) (*Strategy, error) {
	//	change tid into int
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		beego.Error(err)
		return nil, err
	}

	o := orm.NewOrm()

	strategy := new(Strategy)

	qs := o.QueryTable("strategy")
	err = qs.Filter("Id", tidNum).One(strategy)
	if err != nil {
		return nil, err
	}

	return strategy, nil
}

func DeleteStrategy(tid string) error {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		// beego.Error(err)
		return err
	}

	o := orm.NewOrm()
	strategy := &Strategy{Id: tidNum}
	if o.Read(strategy) == nil {
		_, err = o.Delete(strategy)
		if err != nil {
			return err
		}
	}
	return nil
}

func ModifyStrategy(tid string, name string, methods string, conditions string, send_to string) error {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		// beego.Error(err)
		return err
	}

	o := orm.NewOrm()
	strategy := &Strategy{Id: tidNum}
	if o.Read(strategy) == nil {
		strategy.Strategy_name = name
		strategy.Strategy_conditions = conditions
		strategy.Sending_method = methods
		strategy.Send_to = send_to
		_, err = o.Update(strategy)
		if err != nil {
			return err
		}
	}

	return nil
}
