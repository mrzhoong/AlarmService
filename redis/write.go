package redis

import (
	"AlarmService/g"
	"AlarmService/models"
	"log"
)

func ProcData(event models.Event) {
	for _, s := range g.Cfg().Strategys {
		//log.Println("strategys NO:", j)
		for _, v := range s.EquipGroup {
			//log.Println("告警源组:", k)
			if v.Id == event.StrategyID {
				//log.Println("匹配成功", j)
				//	send to queue
				SendToQueue(s.Method, s.To, event)
				//fmt.Println(s.Method, s.To, event.Content)
			} else {
				// fmt.Println("无效事件", j)
				//	send to mysql
			}
		}
	}
}

func SendToQueue(L []g.SendMethod, M []g.SendTo, E models.Event) error {
	for _, v := range L {
		switch v.Name {
		case "短信":
			SendSMS("/sms", M, E)
		case "邮件":
			SendEmail("/mail", M, E)
		case "微信":
			SendWeChat("/wechat", M)
		case "电话":
			SendPhone("/phone", M)
		}
	}
	return nil
}

func SendSMS(queue string, M []g.SendTo, E models.Event) error {
	for _, v := range M {
		log.Println(queue, v.Mobile, E)
		WriteSms(v.Mobile, E)
	}
	return nil
}

func SendEmail(queue string, M []g.SendTo, E models.Event) error {
	for _, v := range M {
		log.Println(queue, v.Email, E.Content)
		WriteMail(v.Email, g.Config().MailTopic, E)
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
