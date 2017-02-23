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
// 	_MYSQL_DRIVER = "MYSQL"
// )

//  策略
type Strategy struct {
	//	唯一性标志
	Id int64
	//	需要发送的设备组编号
	StrategyId         string
	//	策略名称，即组名称
	StrategyName string
	//	策略条件
	StrategyCondition string
	//	发送方式
	SendMethod string
	//	发送人员
	SendTo string
}

//	告警条件
type AlarmConditions struct {
	Id int64
	//	策略ID
	StrategyId int64
	LevelMsg   string
	Content     string
}

//	发送方式
type SendMethods struct {
	Id int64
	//	策略ID
	StrategyId int64
	//	发送方式对应编号
	SendType    []int64
	RepeatInterval int64
}

//	发送方式对应类型
type SendType struct {
	Id       int64
	SendType string
}

// 告警源

// 用户
type User struct {
	//	唯一性标志
	Id int64
	//	用户名称
	Name string
	//	手机号码，非座机
	Tel string
	//	邮箱
	Email string
	//	微信账号
	Wechat string
	//	组织编号
	OrganizationCode int64
	//	组织名称
	OrganizationName string
}

func RegisterDB() {
	dbhost := beego.AppConfig.String("dbhost")
	dbport := beego.AppConfig.String("dbport")
	dbuser := beego.AppConfig.String("dbuser")
	dbpwd := beego.AppConfig.String("dbpwd")
	dbname := beego.AppConfig.String("dbname")
	// 注册模型
	orm.RegisterModel(new(Strategy), new(User), new(SendType))
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
		StrategyName:       name,
		SendMethod:      methods,
		StrategyCondition: conditions,
		SendTo:             send_to,
	}
	_, err := o.Insert(strategy)
	if err != nil {
		return err
	}
	return nil
}

//	获取某个策略
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

//	删除策略
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

//	修改策略
func ModifyStrategy(tid string, name string, methods string, conditions string, send_to string) error {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		// beego.Error(err)
		return err
	}

	o := orm.NewOrm()
	strategy := &Strategy{Id: tidNum}
	if o.Read(strategy) == nil {
		strategy.StrategyName = name
		strategy.StrategyCondition = conditions
		strategy.SendMethod = methods
		strategy.SendTo = send_to
		_, err = o.Update(strategy)
		if err != nil {
			return err
		}
	}

	return nil
}

//	获取策略组信息，组包含设备信息
func GetStrategyGroupInfo() (str string, err error) {
	str = `{"errcode": "00000", "info": [["空调组", "0001"],"UPS组":"0002"]}`

	return str, nil
}

//	获取人员信息列表
func GetUserGroupInfo() (str string, err error) {
	str = `{"errcode": "00000","info": [["小明","000223"],["小敏","000437"]]}`

	return str, nil
}

//	获取Id对应的人名信息
func GetUserName(Id int64) string {
	user := new(User)
	o := orm.NewOrm()
	qs := o.QueryTable("User")
	err := qs.Filter("Id", Id).One(user)
	if err != nil {
		return "未找到"
	}

	return user.Name
}
