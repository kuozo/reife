package main

import (
	"github.com/astaxie/beego"
	ctl "github.com/kollinchu/reife/web/controllers"
)

func main() {
	beego.Router("/", &ctl.HomeController{})
	beego.Router("/about", &ctl.AboutController{})
	beego.Run()
}
