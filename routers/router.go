package routers

import (
	"beerbubble/MtimeCI/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/test", &controllers.TestController{})
	beego.AutoRouter(&controllers.TestController{})
	beego.AutoRouter(&controllers.LoginController{})
	beego.Router("/project", &controllers.ProjectController{})
	beego.AutoRouter(&controllers.ProjectController{})

}
