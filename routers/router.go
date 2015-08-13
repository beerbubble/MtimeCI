package routers

import (
	"beerbubble/MtimeCI/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.TestController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/test", &controllers.TestController{})
	beego.AutoRouter(&controllers.TestController{})
	beego.AutoRouter(&controllers.LoginController{})

}
