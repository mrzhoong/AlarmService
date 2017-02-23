package controllers

import (
	"github.com/astaxie/beego"
	"AlarmService/redis"
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
	EventTime := this.GetString("time")
	Content := this.GetString("content")
	Equip := this.GetString("equip")
	s := EventTime + Content + Equip

	redis.WriteSms(EventTime, Content)

	this.Ctx.WriteString(s)
}
