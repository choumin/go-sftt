package routers

import (
	"zhiwen2/zhiwen2/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
