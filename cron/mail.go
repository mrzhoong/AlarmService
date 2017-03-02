package cron

import (
	"AlarmService/models"
	"AlarmService/redis"
	"encoding/json"
	"log"
)

func WriteMail(tos string, subject, content string) {
	if len(tos) == 0 {
		return
	}

	mail := &models.Mail{Tos: tos, Subject: subject, Content: content}
	WriteMailModel(mail)
}

func WriteMailModel(mail *models.Mail) {
	if mail == nil {
		return
	}

	bs, err := json.Marshal(mail)
	if err != nil {
		log.Println(err)
		return
	}

	redis.LPUSH("/mail", string(bs))
}
