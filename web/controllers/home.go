package controllers

import (
	"github.com/astaxie/beego"
)

type HomeController struct {
	beego.Controller
}

func (this *HomeController) Get() {
	this.Data["Website"] = "Reife"
	this.Data["Email"] = "kollinchu@gmail.com"
	this.TplNames = "home.html"
}
