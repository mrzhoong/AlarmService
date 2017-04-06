package redis

import (
	"AlarmService/models"
	"encoding/json"
	"github.com/astaxie/beego/logs"
	"time"
)

func WriteMail(tos string, subject string, E models.Event) {
	if len(tos) == 0 {
		return
	}
	//	内容
	var mailContent string

	mailContent = time.Unix(E.EventTime, 0).Format("2006-01-02 15:04:05") + " " +
		E.Location + " " + E.EquipName + " " + E.Content

	//	邮件主题显示所有
	if 0 == len(subject) {
		subject = mailContent
	}

	logs.GetLogger("MAIL").Println("subject len:", len(subject))

	mail := &models.Mail{Tos: tos, Subject: subject, Content: mailContent}
	WriteMailModel(mail, E)
}

func WriteMailModel(mail *models.Mail, E models.Event) {
	if mail == nil {
		return
	}

	bs, err := json.Marshal(mail)
	if err != nil {
		logs.GetLogger("MAIL").Println(err)
		return
	}

	LPUSH("/mail", string(bs), E, mail.Tos)
}
