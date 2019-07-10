package routers

import (
	"zhiwen2/zhiwen2/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/paper/:id([0-9]+)", &controllers.PaperController{})
}
