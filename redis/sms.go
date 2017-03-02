package redis

import (
	//"AlarmService/redis"
	"AlarmService/models"
	"encoding/json"
	"log"
)

//type Sms struct {
//	Tos     string `json:"tos"`
//	Content string `json:"content"`
//}

func LPUSH(queue, message string) {
	rc := ConnPool.Get()
	defer rc.Close()
	_, err := rc.Do("LPUSH", queue, message)

	if err != nil {
		log.Println("LPUSH redis", queue, "fail:", err, "message:", message)
	}
}

func WriteSms(tos string, content string) {
	sms := &models.Sms{Tos: tos, Content: content}
	WriteSmsModel(sms)
}

func WriteSmsModel(sms *models.Sms) {
	if sms == nil {
		return
	}

	bs, err := json.Marshal(sms)
	if err != nil {
		log.Println(err)
		return
	}
	LPUSH("/sms", string(bs))
}
