package models

import (
	"AlarmService/g"
	//"AlarmService/redis"
	// "encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"qiniupkg.com/x/errors.v7"
	"strconv"
)

// const (
// 	// 设置数据库路径
// 	_DB_NAME = "data/beeblog.db"
// 	// 设置数据库名称
// 	_MYSQL_DRIVER = "MYSQL"
// )

//	短信
type Sms struct {
	Tos     string `json:"tos"`
	Content string `json:"content"`
}

//	邮件内容
type Mail struct {
	Tos     string `json:"tos"`
	Subject string `json:"subject"`
	Content string `json:"content"`
}

//	发送记录
type SendRecord struct {
	Id int64
	//	站点
	Station string
	//	设备名称
	EquipName string
	//	发生时间
	EventTime int64
	//	事件级别
	EventLevel int64
	//	目标人员
	Tos string
	//	发送方式
	SendType string
	//	发送时间
	SendTime int64
	//	发送状态
	SendStatus bool
	//	发送内容
	Content string
}

//	操作记录

//	告警源分类
type StrategyList struct {
	Id            int64
	Name          string
	StrategyInfos *StrategyInfo `orm:"rel(fk)"`
}

//	发送方式
type SendMethod struct {
	Id            int64
	Name          string
	StrategyInfos *StrategyInfo `orm:"rel(fk)"`
}

//	发送人员信息
type SendTo struct {
	Id            int64
	Name          string
	Mobile        string
	WeChat        string
	Email         string
	StrategyInfos *StrategyInfo `orm:"rel(fk)"`
}

//	策略信息
type StrategyInfo struct {
	//	策略ID
	Id int64
	//	发送的组信息
	EquipGroup []*StrategyList `orm:"reverse(many)"`
	//	策略名称
	Name string
	//	条件，暂未处理
	Condition string
	//	发送方式
	Method []*SendMethod `orm:"reverse(many)"`
	//	发送人员
	To []*SendTo `orm:"reverse(many)"`
}

//	告警事件
type Event struct {
	//	事件时间
	EventTime int64 `json:"eventtime"`
	//	位置信息
	Location string `json:"eventlocation"`
	//	设备ID
	EquipID string `json:"equipid"`
	//	设备名称
	EquipName string `json:"equipname"`
	//	事件级别
	EventLevel int64 `json:"eventlevel"`
	//	事件内容
	Content string `json:"content"`
	//	事件源组
	StrategyID int64 `json:"strategyid"`
	//	是否告警
	IsAlarm bool `json:"isalarm"`
}

//  策略
type Strategy struct {
	//	唯一性标志
	Id int64
	//	需要发送的设备组编号
	StrategyId string
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
	Content    string
}

//	发送方式
type SendMethods struct {
	Id int64
	//	策略ID
	StrategyId int64
	//	发送方式对应编号
	SendType       []int64
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
	// 注册模型
	orm.RegisterModel(new(Strategy), new(User), new(SendType), new(StrategyInfo),
		new(StrategyList), new(SendTo), new(SendMethod), new(SendRecord))
	// mysql 属于默认注册，此处代码可省略）
	orm.RegisterDriver("mysql", orm.DRMySQL)
	// 注册默认数据库
	conn := g.Config().DbUser + ":" + g.Config().DbPwd + "@tcp(" + g.Config().DbHost + ":" +
		g.Config().DbPort + ")/" + g.Config().DbName + "?charset=utf8"
	orm.RegisterDataBase("default", "mysql", conn)
	// orm.RegisterDataBase("default", "mysql", "root:123456@/default?charset=utf8")
}

// 获取策略列表
func GetAllStrategys() ([]*Strategy, error) {
	o := orm.NewOrm()
	strategys := make([]*Strategy, 0)

	qs := o.QueryTable("strategy")
	_, err := qs.All(&strategys)
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
		StrategyName:      name,
		SendMethod:        methods,
		StrategyCondition: conditions,
		SendTo:            send_to,
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
func GetStrategyGroupInfo() (err error) {
	o := orm.NewOrm()
	strategys := new(StrategyInfo)

	qs := o.QueryTable("strategy")
	err = qs.One(strategys)
	if err != nil {
		beego.Error(err)
	}
	beego.Error(strategys.Name)

	return nil
}

//	获取人员信息列表
func GetUserGroupInfo() (str string, err error) {
	str = `{"errcode": "00000","info": [["小明","000223"],["小敏","000437"]]}`
	o := orm.NewOrm()
	k := orm.NewOrm()
	s2 := new(StrategyList)
	s1 := new(StrategyInfo)
	s3 := new(SendTo)
	s4 := new(SendMethod)
	////s1.Id = 2
	s1.EquipGroup = append(s1.EquipGroup, &StrategyList{Name: "暖通"})
	s1.EquipGroup = append(s1.EquipGroup, &StrategyList{Name: "高压"})
	s1.EquipGroup = append(s1.EquipGroup, &StrategyList{Name: "低压"})
	s1.Name = "运维组"
	s1.Condition = "Condition"
	s1.Method = append(s1.Method, &SendMethod{Name: "电话"})
	s1.Method = append(s1.Method, &SendMethod{Name: "微信"})
	s1.To = append(s1.To, &SendTo{Name: "sam2", Mobile: "18503036522", WeChat: "eqwead_ew1_2", Email: "xxx@zktz.com"})

	s2.StrategyInfos = s1
	s2.Name = "暖通"

	s3 = &SendTo{Name: "sam2", Mobile: "18503036522", WeChat: "eqwead_ew1_2", Email: "xxx@zktz.com"}
	s3.StrategyInfos = s1

	s4 = &SendMethod{Name: "电话"}
	s4.StrategyInfos = s1
	//
	_, err = k.Insert(s1)
	if err != nil {
		fmt.Println(err)
		return "nil", err
	}
	_, err = o.Insert(s2)

	if err != nil {
		fmt.Println(err)
		return "nil", err
	}

	_, err = o.Insert(s3)

	if err != nil {
		fmt.Println(err)
		return "nil", err
	}

	_, err = o.Insert(s4)

	if err != nil {
		fmt.Println(err)
		return "nil", err
	}
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

//	发送记录模块
func AddSendRecord(s *SendRecord) error {
	o := orm.NewOrm()

	if s == nil {
		return errors.New("Record is nil")
	}

	_, err := o.Insert(s)
	if err != nil {
		return err
	}
	return nil
}

func GetSendRecord() {

}
