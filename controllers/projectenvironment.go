package controllers

import (
	"beerbubble/MtimeCI/models"
	"beerbubble/MtimeCI/utility"
	//"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	//"strconv"
	//"github.com/astaxie/beego/logs"
)

type ProjectEnvironmentController struct {
	beego.Controller
}

func (this *ProjectEnvironmentController) AddOrEdit() {

	utility.ViewLogin(this.Ctx)

	projectid := this.Input().Get("projectid")
	envid := this.Input().Get("envid")

	var projectenv models.ProjectEnvironment
	var title string

	o := orm.NewOrm()

	if envid == "" {
		title = "添加项目及环境信息"

	} else {
		title = "编辑项目及环境信息"
		//var env models.Environmentinfo
		err := o.QueryTable("Environmentinfo").Filter("envid", envid).Filter("projectid", projectid).One(&projectenv)
		if err == orm.ErrMultiRows {
			// 多条的时候报错
			//fmt.Printf("Returned Multi Rows Not One")
		}
	}

	var envs []*models.Environmentinfo
	num, _ := o.QueryTable("Environmentinfo").All(&envs)
	if num > 0 {

	}

	var project models.Projectinfo
	err := o.QueryTable("Projectinfo").Filter("Id", projectid).One(&project)
	if err == orm.ErrMultiRows || err == orm.ErrNoRows {
	}

	this.Data["Title"] = title
	this.Data["projectenv"] = projectenv
	this.Data["project"] = project
	this.Data["envs"] = envs

	this.Layout = "Template.html"
	this.TplNames = "projectenv/addoredit.html"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["NavContent"] = "component/nav.html"
	this.LayoutSections["Scripts"] = "env/addoreditjs.html"
	this.LayoutSections["HtmlHead"] = "env/listcss.html"
}
