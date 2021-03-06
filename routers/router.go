package routers

import (
	"beerbubble/MtimeCI/controllers"
	"github.com/astaxie/beego"
)

func init() {
	// beego.Router("/", &controllers.ProjectController{}, "*:List")
	beego.Router("/", &controllers.MainController{})

	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/test", &controllers.TestController{})
	beego.AutoRouter(&controllers.TestController{})
	beego.AutoRouter(&controllers.LoginController{})
	beego.Router("/project", &controllers.ProjectController{})
	beego.AutoRouter(&controllers.ProjectController{})
	//beego.Router("/environment", &controllers.EnvironmentController{})
	beego.AutoRouter(&controllers.EnvironmentController{})
	beego.AutoRouter(&controllers.ProjectEnvironmentController{})
	beego.AutoRouter(&controllers.ProjectModuleController{})
	beego.AutoRouter(&controllers.RundeckController{})
}
