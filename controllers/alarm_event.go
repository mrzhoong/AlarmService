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
	EventTime, _ := this.GetInt64("time")
	Content := this.GetString("content")
	StrategyID, _ := this.GetInt64("strategyid")
	Location := this.GetString("location")
	EquipName := this.GetString("equipname")
	Level, _ := this.GetInt64("level")
	IsAlarm, _ := this.GetInt64("isalarm")

	//"2017-03-03 12:00:01 科学城机房 3#温湿度 温度过高报警 当前值：38°C"
	//s := Content
	event := models.Event{EventTime: EventTime, Location: Location, EquipName: EquipName,
		Content: Content, EquipID: "00245",
		EventLevel: Level, StrategyID: StrategyID, IsAlarm: IsAlarm}
	redis.ProcData(event)
	this.Ctx.WriteString("OK")
}

func (c *EventController) Post() {

	c.Ctx.Output.Header("Access-Control-Allow-Origin", "*")
	c.Ctx.Output.Header("Access-Control-Allow-Headers", "X-Requested-With")
	c.Ctx.Output.Header("Access-Control-Allow-Method", "POST,GET,OPTIONS")

	var ob models.Event
	json.Unmarshal(c.Ctx.Input.RequestBody, &ob)

	redis.ProcData(ob)

	ret := "{\n"
	ret += "\"errcode\": \"00001\",\n"
	ret += "\"errmsg\": \"OK\"\n}\n"

	c.Ctx.WriteString(ret)
	return
}
