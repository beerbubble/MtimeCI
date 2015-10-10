package controllers

import (
	"beerbubble/MtimeCI/datatype"
	"beerbubble/MtimeCI/models"
	"beerbubble/MtimeCI/utility"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	//"github.com/astaxie/beego/logs"
	// "strconv"
	//"bytes"
	"fmt"
	// "os/exec"
	// "strings"
)

type ProjectModuleController struct {
	beego.Controller
}

func (this *ProjectModuleController) AddOrEdit() {

	utility.ViewLogin(this.Ctx)

	envid := this.Input().Get("envid")

	var env models.Environmentinfo
	var title string

	//envTypes := datatype.EnvTypes

	for key, value := range datatype.EnvTypeMap {
		fmt.Println(key)
		fmt.Println(value)
	}

	if envid == "" {
		title = "添加项目模块信息"

	} else {
		title = "编辑项目模块信息"

		o := orm.NewOrm()

		//var env models.Environmentinfo
		o.QueryTable("Environmentinfo").Filter("id", envid).One(&env)
		//if err == orm.ErrMultiRows {
		// 多条的时候报错
		//fmt.Printf("Returned Multi Rows Not One")
		//}
	}

	this.Data["Title"] = title
	this.Data["env"] = env

	this.Layout = "Template.html"
	this.TplNames = "projectmodule/addoredit.html"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["NavContent"] = "component/nav.html"
	this.LayoutSections["Scripts"] = "projectmodule/addoreditjs.html"
	this.LayoutSections["HtmlHead"] = "env/listcss.html"
}
