package redis

import (
	"AlarmService/g"
	"log"
)

//type StrategyList struct {
//	Id   int64  `json:id`
//	Name string `json:name`
//}
//
//type SendMethod struct {
//	Id   int64  `json:id`
//	Name string `json:name`
//}
//
//type SendTo struct {
//	Id     int64  `json:id`
//	Name   string `json:name`
//	Mobile string `json:mobile`
//	WeChat string `json:wechat`
//	Email  string `json:email`
//}

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

//
//type StrategyInfo struct {
//	//	策略ID
//	Id int64 `json:id`
//	//	发送的组信息
//	EquipGroup []StrategyList `json:strategylist`
//	//	策略名称
//	Name string `json:name`
//	//	条件，暂未处理
//	Condition string `json:condition`
//	//	发送方式
//	Method []SendMethod `json:sendmethod`
//	//	发送人员
//	To []SendTo `json:sendto`
//}

func ProcData(event Event) {
	for j, s := range g.Cfg().Strategys {
		log.Println("strategys NO:", j)
		for k, v := range s.EquipGroup {
			log.Println("告警源组:", k)
			if v.Id == event.StrategyID {
				log.Println("匹配成功", j)
				//	send to queue
				SendToQueue(s.Method, s.To, event.Content)
			} else {
				// fmt.Println("无效事件", j)
				//	send to mysql
			}
		}
	}
}

func SendToQueue(L []g.SendMethod, M []g.SendTo, Content string) error {
	for _, v := range L {
		switch v.Name {
		case "短信":
			SendSMS("/sms", M, Content)
		case "邮件":
			SendEmail("/mail", M, Content)
		case "微信":
			SendWeChat("/wechat", M)
		case "电话":
			SendPhone("/phone", M)
		}
	}
	return nil
}

func SendSMS(queue string, M []g.SendTo, Content string) error {
	for _, v := range M {
		log.Println(queue, v.Mobile, Content)
		WriteSms(v.Mobile, Content)
	}
	return nil
}

func SendEmail(queue string, M []g.SendTo, Content string) error {
	for _, v := range M {
		log.Println(queue, v.Email, Content)
		WriteMail(v.Email, g.Config().MailTopic, Content)
	}
	return nil
}

func SendWeChat(queue string, M []g.SendTo) error {
	for k, v := range M {
		log.Println(queue, v.WeChat, "No:", k)
	}
	return nil
}

func SendPhone(queue string, M []g.SendTo) error {
	for k, v := range M {
		log.Println(queue, v.Mobile, "No:", k)
	}
	return nil
}

func AddStrategy() error {
	return nil
}
