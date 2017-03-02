package redis

import (
	"fmt"
	//"github.com/astaxie/beego/orm"
	//_ "github.com/go-sql-driver/mysql"
	//"github.com/astaxie/beego/orm"
	//"github.com/astaxie/beego/orm"
)

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

type Event struct {
	//	事件时间
	EventTime int64 `json:"eventtime"`
	//	位置信息
	EventLocation string `json:"eventlocation"`
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

func ProcData() {
	m := make(map[int64]StrategyInfo)
	var s StrategyInfo
	s.Id = 1
	s.EquipGroup = append(s.EquipGroup, StrategyList{Id: 1, Name: "name"})
	s.EquipGroup = append(s.EquipGroup, StrategyList{Id: 2, Name: "name"})
	s.EquipGroup = append(s.EquipGroup, StrategyList{Id: 3, Name: "暖通"})
	s.Name = "运维组"
	s.Condition = "Condition"
	s.Method = append(s.Method, SendMethod{Id: 1, Name: "短信"})
	s.To = append(s.To, SendTo{Id: 1, Name: "sam", Mobile: "18503036521", WeChat: "eqwead_ew1_2", Email: "xxx@zktz.com"})
	m[1] = s

	var s1 StrategyInfo
	s1.Id = 2
	s1.EquipGroup = append(s1.EquipGroup, StrategyList{Id: 1, Name: "高压2"})
	s1.EquipGroup = append(s1.EquipGroup, StrategyList{Id: 2, Name: "抵押2"})
	s1.EquipGroup = append(s1.EquipGroup, StrategyList{Id: 3, Name: "暖通2"})
	s1.Name = "运维组"
	s1.Condition = "Condition"
	s1.Method = append(s1.Method, SendMethod{Id: 2, Name: "电话"})
	s1.Method = append(s1.Method, SendMethod{Id: 3, Name: "微信"})
	s1.To = append(s1.To, SendTo{Id: 1, Name: "sam2", Mobile: "18503036522", WeChat: "eqwead_ew1_2", Email: "xxx@zktz.com"})
	m[2] = s1

	var s2 StrategyInfo
	s2.Id = 3
	s2.EquipGroup = append(s2.EquipGroup, StrategyList{Id: 1, Name: "风冷3"})
	s2.EquipGroup = append(s2.EquipGroup, StrategyList{Id: 2, Name: "水冷3"})
	s2.EquipGroup = append(s2.EquipGroup, StrategyList{Id: 3, Name: "弱点3"})
	s2.Name = "运维组"
	s2.Condition = "Condition"
	s2.Method = append(s2.Method, SendMethod{Id: 3, Name: "微信"})
	s2.To = append(s2.To, SendTo{Id: 1, Name: "sam3", Mobile: "18503036523", WeChat: "eqwead_ew1_2", Email: "xxx@zktz.com"})
	m[3] = s2

	var s3 StrategyInfo
	s3.Id = 4
	s3.EquipGroup = append(s3.EquipGroup, StrategyList{Id: 1, Name: "name4"})
	s3.EquipGroup = append(s3.EquipGroup, StrategyList{Id: 2, Name: "name4"})
	s3.EquipGroup = append(s3.EquipGroup, StrategyList{Id: 3, Name: "暖通4"})
	s3.Name = "运维组"
	s3.Condition = "Condition"
	s3.Method = append(s3.Method, SendMethod{Id: 4, Name: "邮件"})
	s3.To = append(s3.To, SendTo{Id: 1, Name: "sam4", Mobile: "18503036524", WeChat: "eqwead_ew1_2", Email: "xxx@zktz.com"})
	m[4] = s3
	//	2.Event
	event := Event{EventTime: 1488092446, EventLocation: "科学城机房", EquipName: "3#温湿度", Content: "湿度过高报警",
		EquipID: "00245", EventLevel: 3, StrategyID: 3, IsAlarm: false}
	// fmt.Println(event.Content)
	//	3.If
	for j, s := range m {
		d := s.EquipGroup
		for _, v := range d {
			if v.Id == event.StrategyID {
				fmt.Println("匹配成功", j)
				//	send to queue
				SendToQueue(s.Method, s.To)
			} else {
				// fmt.Println("无效事件", j)
				//	send to mysql
			}
		}
	}

}

func SendToQueue(L []SendMethod, M []SendTo) error {
	for _, v := range L {
		switch v.Name {
		case "短信":
			SendSMS("/sms", M)
		case "邮件":
			SendEmail("/mail", M)
		case "微信":
			SendWeChat("/wechat", M)
		case "电话":
			SendPhone("/phone", M)
		}
	}
	return nil
}

func SendSMS(queue string, M []SendTo) error {
	for k, v := range M {
		fmt.Println(queue, v.Mobile, "No:", k)
	}
	return nil
}

func SendEmail(queue string, M []SendTo) error {
	for k, v := range M {
		fmt.Println(queue, v.Email, "No:", k)
	}
	return nil
}

func SendWeChat(queue string, M []SendTo) error {
	for k, v := range M {
		fmt.Println(queue, v.WeChat, "No:", k)
	}
	return nil
}

func SendPhone(queue string, M []SendTo) error {
	for k, v := range M {
		fmt.Println(queue, v.Mobile, "No:", k)
	}
	return nil
}

func AddStrategy() error {
	//o := orm.NewOrm()
	//var s1 StrategyInfo
	//s1.Id = 2
	//s1.EquipGroup = append(s1.EquipGroup, StrategyList{Id: 1, Name: "高压2"})
	//s1.EquipGroup = append(s1.EquipGroup, StrategyList{Id: 2, Name: "抵押2"})
	//s1.EquipGroup = append(s1.EquipGroup, StrategyList{Id: 3, Name: "暖通2"})
	//s1.Name = "运维组"
	//s1.Condition = "Condition"
	//s1.Method = append(s1.Method, SendMethod{Id: 2, Name: "电话"})
	//s1.Method = append(s1.Method, SendMethod{Id: 3, Name: "微信"})
	//s1.To = append(s1.To, SendTo{Id: 1, Name: "sam2", Mobile: "18503036522", WeChat: "eqwead_ew1_2", Email: "xxx@zktz.com"})
	//_, err := o.Insert(s1)
	//
	//if err != nil {
	//	return err
	//}
	return nil
}
