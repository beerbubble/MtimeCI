package controllers

import (
	"beerbubble/MtimeCI/models"
	"beerbubble/MtimeCI/utility"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
	//"github.com/astaxie/beego/logs"
)

type ProjectEnvironmentController struct {
	beego.Controller
}

func (this *ProjectEnvironmentController) AddOrEdit() {

	utility.ViewLogin(this.Ctx)

	projectid := this.Input().Get("projectid")
	envid := this.Input().Get("envid")

	var projectenv models.Projectenvironment
	var title string

	o := orm.NewOrm()

	if envid == "" {
		title = "添加项目及环境信息"

	} else {
		title = "编辑项目及环境信息"
		//var env models.Environmentinfo
		err := o.QueryTable("Projectenvironment").Filter("envid", envid).Filter("projectid", projectid).One(&projectenv)
		if err == orm.ErrMultiRows {
			// 多条的时候报错
			//fmt.Printf("Returned Multi Rows Not One")
		}
	}

	var envs []*models.Environmentinfo
	num, _ := o.QueryTable("Environmentinfo").All(&envs)
	if num > 0 {

	}

	viewenvironmentinfomodels := []*models.ViewEnvironmentinfoModel{}
	for i := 0; i < len(envs); i++ {
		selected := ""

		if envs[i].Id == projectenv.Envid {
			selected = "selected"
		}

		viewenvironmentinfomodels = append(viewenvironmentinfomodels, &models.ViewEnvironmentinfoModel{envs[i].Id, envs[i].Name, envs[i].Description, envs[i].Rundeckapiurl, selected})
	}

	var project models.Projectinfo
	err := o.QueryTable("Projectinfo").Filter("Id", projectid).One(&project)
	if err == orm.ErrMultiRows || err == orm.ErrNoRows {
	}

	this.Data["Title"] = title
	this.Data["projectenv"] = projectenv
	this.Data["project"] = project
	this.Data["envs"] = viewenvironmentinfomodels

	this.Layout = "Template.html"
	this.TplNames = "projectenv/addoredit.html"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["NavContent"] = "component/nav.html"
	this.LayoutSections["Scripts"] = "projectenv/addoreditjs.html"
	this.LayoutSections["HtmlHead"] = "env/listcss.html"
}

func (this *ProjectEnvironmentController) AddApi() {

	envid := this.Input().Get("envid")
	projectenvid := this.Input().Get("projectenvid")
	rundeckbuildjobid := this.Input().Get("rundeckbuildjobid")
	rundeckpackagejobid := this.Input().Get("rundeckpackagejobid")
	projectid := this.Input().Get("projectid")

	fmt.Println(rundeckbuildjobid)

	intenvid, _ := strconv.Atoi(envid)
	intprojectenvid, _ := strconv.Atoi(projectenvid)
	intprojectid, _ := strconv.Atoi(projectid)

	o := orm.NewOrm()

	if intprojectenvid > 0 {

		projectenv := models.Projectenvironment{Id: intprojectenvid}
		if o.Read(&projectenv) == nil {
			projectenv.Projectid = intprojectid
			projectenv.Envid = intenvid
			projectenv.Rundeckbuildjobid = rundeckbuildjobid
			projectenv.Rundeckpackagejobid = rundeckpackagejobid
			if num, err := o.Update(&projectenv); err == nil {
				this.Data["json"] = models.EnvAddModel{models.JsonResultBaseStruct{Result: true, Message: "操作成功"}, num}
				this.ServeJson()
			}
		}

	} else {
		var projectenv models.Projectenvironment
		err := o.QueryTable("Projectenvironment").Filter("envid", envid).Filter("projectid", projectid).One(&projectenv)
		if err == orm.ErrMultiRows || err == orm.ErrNoRows || projectenv.Id > 0 {
			this.Data["json"] = models.JsonResultBaseStruct{Result: false, Message: "添加重复数据!"}
			this.ServeJson()
		} else {
			var projectenv models.Projectenvironment
			projectenv.Projectid = intprojectid
			projectenv.Envid = intenvid
			projectenv.Rundeckbuildjobid = rundeckbuildjobid
			projectenv.Rundeckpackagejobid = rundeckpackagejobid

			if id, err := o.Insert(&projectenv); err == nil {
				this.Data["json"] = models.EnvAddModel{models.JsonResultBaseStruct{Result: true, Message: "操作成功"}, id}
				this.ServeJson()
			}
		}
	}
}

func (this *ProjectEnvironmentController) Build() {

	utility.ViewLogin(this.Ctx)

	projectid := this.Input().Get("projectid")
	envid := this.Input().Get("envid")

	var projectenv models.Projectenvironment
	var env *models.Environmentinfo
	var title string

	o := orm.NewOrm()

	var envs []*models.Environmentinfo
	o.QueryTable("Environmentinfo").All(&envs)

	o.QueryTable("Projectenvironment").Filter("envid", envid).Filter("projectid", projectid).One(&projectenv)

	viewenvironmentinfomodels := []*models.ViewEnvironmentinfoModel{}
	for i := 0; i < len(envs); i++ {
		selected := ""

		if envs[i].Id == projectenv.Envid {
			selected = "selected"
			env = envs[i]
		}

		viewenvironmentinfomodels = append(viewenvironmentinfomodels, &models.ViewEnvironmentinfoModel{envs[i].Id, envs[i].Name, envs[i].Description, envs[i].Rundeckapiurl, selected})
	}

	var project models.Projectinfo
	err := o.QueryTable("Projectinfo").Filter("Id", projectid).One(&project)
	if err == orm.ErrMultiRows || err == orm.ErrNoRows {
	}

	title = project.Name + " " + env.Name + " 部署详情"

	var projectbranchs []*models.Projectbranch
	o.QueryTable("Projectbranch").Filter("Projectid", projectid).All(&projectbranchs)

	this.Data["Title"] = title
	this.Data["projectenv"] = projectenv
	this.Data["project"] = project
	this.Data["envs"] = viewenvironmentinfomodels
	this.Data["envid"] = envid
	this.Data["projectbranchs"] = projectbranchs

	this.Layout = "Template.html"
	this.TplNames = "projectenv/build.html"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["NavContent"] = "component/nav.html"
	this.LayoutSections["Scripts"] = "projectenv/addoreditjs.html"
	this.LayoutSections["HtmlHead"] = "env/listcss.html"
}
