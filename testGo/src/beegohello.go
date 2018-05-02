package main

import (
	"github.com/astaxie/beego"
)

type HomeController struct {
	beego.Controller
}

func (this *HomeController) Get() {
	this.Ctx.WriteString("hello")
}

func main() {
	beego.Router("/h", &HomeController{})
	beego.Run()
}
