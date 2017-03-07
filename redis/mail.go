package redis

import (
	"AlarmService/models"
	"encoding/json"
	"log"
)

func WriteMail(tos string, subject string, E models.Event) {
	if len(tos) == 0 {
		return
	}

	mail := &models.Mail{Tos: tos, Subject: subject, Content: E.Content}
	WriteMailModel(mail, E)
}

func WriteMailModel(mail *models.Mail, E models.Event) {
	if mail == nil {
		return
	}

	bs, err := json.Marshal(mail)
	if err != nil {
		log.Println(err)
		return
	}

	LPUSH("/mail", string(bs), E)
}
