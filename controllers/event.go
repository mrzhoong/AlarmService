package controllers

import (
	"github.com/astaxie/beego"
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
	EquipId := this.GetString("equip")
	s := EventTime + Content + EquipId
	this.Ctx.WriteString(s)
}
