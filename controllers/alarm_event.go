package controllers

import (
	"AlarmService/models"
	"AlarmService/redis"
	"github.com/astaxie/beego"
	// "log"
	"encoding/json"
	"fmt"
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
	Location := this.GetString("location")
	EquipName := this.GetString("equipname")
	Level, _ := this.GetInt64("level")
	IsAlarm, _ := this.GetBool("isalarm")

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

	//url := c.Ctx.Request.RequestURI
	//	fmt.Println("[" + url + "]")
	fmt.Println(c.Ctx.Request.RequestURI)
	//fmt.Println(string(c.Ctx.Input.RequestBody))

	var ob models.Event
	json.Unmarshal(c.Ctx.Input.RequestBody, &ob)

	fmt.Println(ob.Content)
	c.Ctx.WriteString(string(c.Ctx.Input.RequestBody))
	return
}
