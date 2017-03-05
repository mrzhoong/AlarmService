package controllers

import (
	"AlarmService/models"
	"AlarmService/redis"
	"github.com/astaxie/beego"
	// "log"
	"encoding/json"
)

type EventController struct {
	beego.Controller
}

func (this *EventController) Get() {
	// c.TplName = "event.html"
	// EventTime int64
	// Level int64
	// Location string
	// Equip string
	// EquipId string
	// Content string
	// url:= this.URLFor(endpoint, ...)
	EventTime, _ := this.GetInt64("time")
	Content := this.GetString("content")
	StrategyID, _ := this.GetInt64("strategyid")

	//"2017-03-03 12:00:01 科学城机房 3#温湿度 温度过高报警 当前值：38°C"
	s := Content
	event := models.Event{EventTime: EventTime, Location: "科学城机房", EquipName: "3#温湿度",
		Content: Content, EquipID: "00245",
		EventLevel: 4, StrategyID: StrategyID, IsAlarm: false}
	redis.ProcData(event)
	this.Ctx.WriteString(s)
}

func (this *EventController) Post() {
	var event models.Event
	json.Unmarshal(this.Ctx.Input.RequestBody, &event)
	// this.Ctx.WriteString(event.EquipID)
}
