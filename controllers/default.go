package controllers

import (
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
)

type MainController struct {
	web.Controller
}

func (this *MainController) Get() {
	logs.Info("1")
	this.Ctx.WriteString("hello get")
}

func (this *MainController) Post() {
	logs.Info("1")
	this.Ctx.WriteString("hello post")
}
