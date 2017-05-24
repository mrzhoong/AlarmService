package redis

import (
	"encoding/json"
	"log"
	"time"

	"AlarmService/models"
)

func LPUSH(queue, message string, E models.Event, tos string) {
	rc := ConnPool.Get()
	defer rc.Close()
	_, err := rc.Do("LPUSH", queue, message)

	var R models.SendRecord
	R.Content = E.Content
	R.EquipName = E.EquipName
	R.EventLevel = E.EventLevel
	R.EventTime = E.EventTime
	R.SendTime = time.Now().Unix()
	R.SendType = queue
	R.Station = E.Location
	R.Tos = tos

	if err != nil {
		log.Println("LPUSH redis", queue, "fail:", err, "message:", message)
		R.SendStatus = false
		models.AddSendRecord(&R)
	} else {
		R.SendStatus = true
		models.AddSendRecord(&R)
	}
}

func WriteSms(tos string, E models.Event) {
	sms := &models.Sms{Tos: tos, Content: E.Content}

	WriteSmsModel(sms, E)
}

func WriteSmsModel(sms *models.Sms, E models.Event) {
	if sms == nil {
		return
	}

	bs, err := json.Marshal(sms)
	if err != nil {
		log.Println(err)
		return
	}
	LPUSH("/sms", string(bs), E, sms.Tos)
}
