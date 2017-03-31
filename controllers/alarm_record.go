package controllers

import (
	"AlarmService/models"
	"github.com/astaxie/beego"
	"time"
)

type RecordController struct {
	beego.Controller
}

func (c *RecordController) Get() {
	c.TplName = "eventlog.html"
	beginTime := time.Now().Format("2006-01-02 ") + "00:00:00"
	endTime := time.Now().Format("2006-01-02 15:04:06")
	c.Data["begin_time"] = beginTime
	c.Data["end_time"] = endTime
}

func (c *RecordController) Post() {
	c.TplName = "eventlog.html"

	beginTime := c.Input().Get("begin_time")
	endTime := c.Input().Get("end_time")
	location := c.Input().Get("station")
	sendTo := c.Input().Get("send_to")
	equipName := c.Input().Get("equip_name")
	sendType := c.Input().Get("send_type")

	records, err := models.GetSendRecord(beginTime, endTime, location, equipName, sendTo, sendType)

	if err != nil {
		beego.Error(err)
	}

	c.Data["SendRecord"] = records
	c.Data["begin_time"] = beginTime
	c.Data["end_time"] = endTime
	c.Data["send_type"] = sendType
	c.Data["equip_name"] = equipName
	c.Data["location"] = location
	c.Data["send_to"] = sendTo
}
