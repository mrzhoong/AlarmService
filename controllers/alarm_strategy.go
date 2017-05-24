package controllers

import (
	"AlarmService/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type StrategyController struct {
	beego.Controller
}

func (this *StrategyController) Get() {
	this.Data["isStrategy"] = true
	this.TplName = "strategy.html"
	//	策略列表
	strategys, err := models.GetAllStrategys()

	if err != nil {
		beego.Error(err)
	}
	// this.Data["Strategy_id"] = 1
	// this.Data["Strategy_name"] = "科技部"
	// this.Data["Strategy_conditions"] = "实时"
	// this.Data["Sending_method"] = "短信"
	// this.Data["Send_to"] = "张三，李四"
	//str := []Strategy{Strategy{Strategy_id: 1, Strategy_name: "联通", Strategy_conditions: "实时", Sending_method: "微信", Send_to: "秦羽"},
	//	Strategy{Strategy_id: 2, Strategy_name: "华为", Strategy_conditions: "实时", Sending_method: "短信", Send_to: "张三"},
	//	Strategy{Strategy_id: 3, Strategy_name: "科技部", Strategy_conditions: "定时", Sending_method: "邮件", Send_to: "李四,王五"},
	//	Strategy{Strategy_id: 4, Strategy_name: "中联值班组", Strategy_conditions: "实时", Sending_method: "电话", Send_to: "332"}}
	this.Data["AlarmStrategy"] = strategys
	return
}

func (this *StrategyController) Post() {
	// exchange json
	// var ob models.Object
	// json.Unmarshal(this.Ctx.Input.RequestBody, &ob)
	// is login  get token
	this.Data["IsLogin"] = true
	//	this.TplName = "strategy.html"
	tid := this.Input().Get("tid")
	name := this.Input().Get("strategy_name")
	methods := this.Input().Get("sending_method")
	conditions := this.Input().Get("strategy_condition")
	send_to := this.Input().Get("send_to")
	var err error
	if len(tid) == 0 {
		err = models.AddStrategy(name, methods, conditions, send_to)
		logs.Debug(send_to)
	} else {
		err = models.ModifyStrategy(tid, name, methods, conditions, send_to)
		logs.Info(send_to)
	}

	if err != nil {
		beego.Error(err)
	}
	this.Redirect("/strategy", 302)
}

func (this *StrategyController) Add() {
	// check token

	this.TplName = "strategy_add.html"

	this.Data["User"] = []string{"过敏", "张三", "李四", "王五", "赵六"}
	this.Data["SendType"] = []string{"短信", "电话", "微信", "邮件"}

	return
}

func (this *StrategyController) Delete() {
	// check login

	err := models.DeleteStrategy(this.Input().Get("tid"))
	if err != nil {
		beego.Error(err)
	}

	this.Redirect("/strategy", 302)
}

func (this *StrategyController) Modify() {
	// check login
	this.TplName = "strategy_modify.html"
	tid := this.Input().Get("tid")

	strategy, err := models.GetStrategy(tid)
	//strategy, tos, sType, err := models.GetStrategyConf(tid)
	if err != nil {
		beego.Error(err)
		this.Redirect("/strategy", 302)
		return
	}

	this.Data["Strategy"] = strategy
	this.Data["Tid"] = tid
	this.Data["IsLogin"] = true

	this.Data["User"] = []string{"过敏", "张三", "李四", "王五", "赵六"}
	this.Data["SendType"] = []string{"短信", "电话", "微信", "邮件"}
	// return
}

func (this *StrategyController) View() {
	// check login
	// models.GetUserGroupInfo()
	models.GetStrategyGroupInfo()
	this.Ctx.WriteString("asd")
	return
}
