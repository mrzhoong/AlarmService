package controllers

import (
	"AlarmService/models"
	"github.com/astaxie/beego"
	"log"
	"strconv"
)

type TestController struct {
	beego.Controller
}

func (c *TestController) Get() {
	//c.TplName = "home.html"
	id := c.Input().Get("id")
	log.Println(id)
	d, _ := strconv.Atoi(id)
	//str, _ := models.GetUserGroupInfo()
	str := models.GetUserName(d)
	c.Ctx.WriteString(str)
}
