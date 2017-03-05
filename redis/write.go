package redis

import (
	"AlarmService/g"
	"AlarmService/models"
	"log"
)

func ProcData(event models.Event) {
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
