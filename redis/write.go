package redis

import (
	"AlarmService/g"
	"AlarmService/models"

	"github.com/astaxie/beego/logs"
)

func ProcData(event models.Event) {
	for _, s := range g.Cfg().Strategys {
		for _, v := range s.EquipGroup {
			//log.Println("告警源组:", k)
			if v.Id == event.StrategyID {
				//	send to queue
				SendToQueue(s.Method, s.To, event)
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
			SendWeChat("/wechat", M, E)
		case "电话":
			SendPhone("/phone", M, E)
		}
	}
	return nil
}

func SendSMS(queue string, M []g.SendTo, E models.Event) error {
	for _, v := range M {
		logs.GetLogger("SMS").Println(queue, v.Mobile, E)
		WriteSms(v.Mobile, E)
	}
	return nil
}

func SendEmail(queue string, M []g.SendTo, E models.Event) error {
	for _, v := range M {
		logs.GetLogger("MAIL").Println(queue, v.Email, E.Content)
		WriteMail(v.Email, g.Config().MailTopic, E)
	}
	return nil
}

func SendWeChat(queue string, M []g.SendTo, E models.Event) error {
	for k, v := range M {
		logs.GetLogger("WECHAT").Println(queue, v.WeChat, "No:", k)
	}
	return nil
}

func SendPhone(queue string, M []g.SendTo, E models.Event) error {
	for k, v := range M {
		logs.GetLogger("PHONE").Println(queue, v.Mobile, "No:", k)
		WritePhone(v.Mobile, E)
	}
	return nil
}
