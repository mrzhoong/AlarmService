package controllers

import (
	"AlarmService/models"
	"github.com/astaxie/beego"
)

type RecordController struct {
	beego.Controller
}

func (c *RecordController) Get() {
	c.TplName = "eventlog.html"
}

func (c *RecordController) Post() {
	c.TplName = "eventlog.html"
	var L []models.SendRecord
	s := models.SendRecord{Station: "科学城机房", EquipName: "佳力图空调", EventTime: 165854,
		EventLevel: 1, Tos: "过敏", SendType: "短信", SendTime: 15555, SendStatus: true, Content: "关机"}
	L = append(L, s)
	L = append(L, s)
	L = append(L, s)
	L = append(L, s)
	L = append(L, s)
	L = append(L, s)
	L = append(L, s)
	L = append(L, s)
	c.Data["SendRecord"] = L
}
