package controllers

import (
	//"beerbubble/MtimeCI/datatype"
	"beerbubble/MtimeCI/utility"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	//"strconv"
)

type TestController struct {
	beego.Controller
}

func (this *TestController) Get() {
	//utility.Login(this.Ctx)
	//this.Ctx.SetCookie("MtimeCIUserId", "1001", 180*60, "/")

	this.Layout = "Template.html"
	this.TplNames = "test/index.html"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["NavContent"] = "component/nav.html"
}

func (this *TestController) Jia() {

	utility.ViewLogin(this.Ctx)

	log := logs.NewLogger(10000)
	log.SetLogger("console", "")

	//log.Info(strconv.Itoa(int(datatype.Golang)))

	log.Info(this.Ctx.Input.Uri())
	this.TplNames = "test/index.html"
}
