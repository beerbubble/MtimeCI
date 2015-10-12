package controllers

import (
	"beerbubble/MtimeCI/datatype"
	"beerbubble/MtimeCI/models"
	"beerbubble/MtimeCI/utility"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	//"github.com/astaxie/beego/logs"
	"strconv"
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

	projectid := this.Input().Get("projectid")
	moduleid := this.Input().Get("moduleid")

	o := orm.NewOrm()

	var project models.Projectinfo
	o.QueryTable("Projectinfo").Filter("Id", projectid).One(&project)

	var module models.Moduleinfo

	var title string

	//envTypes := datatype.EnvTypes

	for key, value := range datatype.EnvTypeMap {
		fmt.Println(key)
		fmt.Println(value)
	}

	if moduleid == "" {
		title = "添加项目模块信息"

	} else {
		title = "编辑项目模块信息"

		//var env models.Environmentinfo
		o.QueryTable("Moduleinfo").Filter("id", moduleid).One(&module)
		//if err == orm.ErrMultiRows {
		// 多条的时候报错
		//fmt.Printf("Returned Multi Rows Not One")
		//}
	}

	this.Data["Title"] = title
	// this.Data["env"] = env
	this.Data["project"] = project
	this.Data["module"] = module

	this.Layout = "Template.html"
	this.TplNames = "projectmodule/addoredit.html"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["NavContent"] = "component/nav.html"
	this.LayoutSections["Scripts"] = "projectmodule/addoreditjs.html"
	this.LayoutSections["HtmlHead"] = "env/listcss.html"
}

func (this *ProjectModuleController) AddApi() {

	moduleid, _ := strconv.Atoi(this.Input().Get("moduleid"))
	projectid, _ := strconv.Atoi(this.Input().Get("projectid"))
	modulename := this.Input().Get("modulename")
	moduledes := this.Input().Get("moduledes")

	fmt.Println(modulename)

	o := orm.NewOrm()

	if moduleid > 0 {
		module := models.Moduleinfo{Id: moduleid}
		if o.Read(&module) == nil {
			module.Name = modulename
			module.Description = moduledes
			module.Projectid = projectid
			if num, err := o.Update(&module); err == nil {
				this.Data["json"] = models.EnvAddModel{models.JsonResultBaseStruct{Result: true, Message: "操作成功"}, num}
			}
		}

	} else {
		var module models.Moduleinfo
		module.Name = modulename
		module.Description = moduledes
		module.Projectid = projectid
		if id, err := o.Insert(&module); err == nil {
			this.Data["json"] = models.EnvAddModel{models.JsonResultBaseStruct{Result: true, Message: "操作成功"}, id}
		}
	}

	this.ServeJson()
}

func (this *ProjectModuleController) List() {

	utility.ViewLogin(this.Ctx)

	projectid := this.Ctx.Input.Params["0"]

	o := orm.NewOrm()

	var modules []*models.Moduleinfo
	o.QueryTable("Moduleinfo").Filter("Projectid", projectid).All(&modules)

	this.Data["Title"] = "模块列表"
	//this.Data["envsnum"] = num
	// this.Data["envTypeList"] = envTypeList
	this.Data["modules"] = modules

	this.Layout = "Template.html"
	this.TplNames = "projectmodule/list.html"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["NavContent"] = "component/nav.html"
	this.LayoutSections["Scripts"] = "env/listjs.html"
	this.LayoutSections["HtmlHead"] = "env/listcss.html"
}
