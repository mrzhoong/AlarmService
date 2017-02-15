package controllers

import (
	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

func (this *UserController) Get() {
	this.Data["IsUser"] = true
	this.TplName = "login.html"
	return
}

func (this *UserController) Add() {
	this.Ctx.WriteString("hahahaha")
	return
}
