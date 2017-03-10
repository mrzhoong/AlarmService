package controllers

import (
	"github.com/astaxie/beego"
)

type OperationController struct {
	beego.Controller
}

func (c *OperationController) Get() {
	c.TplName = "operation.html"
	//c.Ctx.WriteString("Coding")
}
