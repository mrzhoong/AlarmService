package controllers

import (
	//"AlarmService/models"
	"github.com/astaxie/beego"
)

type TestController struct {
	beego.Controller
}

func (c *TestController) Get() {
	//c.TplName = "home.html"
	//str, _ := models.GetStrategyGroupInfo()
	c.Ctx.WriteString("test")
}
