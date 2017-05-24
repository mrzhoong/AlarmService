package redis

import (
	"encoding/json"
	"time"

	"AlarmService/models"

	"github.com/astaxie/beego/logs"
)

func WritePhone(tos string, E models.Event) {
	if len(tos) == 0 {
		return
	}
	//	内容
	var phoneContent string

	phoneContent = time.Unix(E.EventTime, 0).Format("2006-01-02 15:04:05") + " " +
		E.Location + " " + E.EquipName + " " + E.Content

	phone := &models.Phone{Guid: models.GetGuid(), Tos: tos, Content: phoneContent}
	WritePhoneModel(phone, E)
}

func WritePhoneModel(phone *models.Phone, E models.Event) {
	if phone == nil {
		return
	}

	bs, err := json.Marshal(E)
	if err != nil {
		logs.GetLogger("PHONE").Println(err)
		return
	}

	LPUSH("/phone", string(bs), E, phone.Tos, phone.Guid)
}
