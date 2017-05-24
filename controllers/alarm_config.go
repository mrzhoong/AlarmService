package controllers

import (
	"AlarmService/g"

	"github.com/astaxie/beego"
)

type ConfigController struct {
	beego.Controller
}

func (c *ConfigController) Get() {
	//c.Ctx.WriteString("Coding")
	c.TplName = "setting.html"
}

func (c *ConfigController) Reload() {

	g.LoadConfig()
	g.LoadCfg("D:\\Development\\src\\AlarmService\\cfg.json")

	ret := "{\n"
	ret += "\"errcode\": \"00000\",\n"
	ret += "\"errmsg\": \"Reload done.\"\n}\n"

	c.Ctx.WriteString(ret)
}
