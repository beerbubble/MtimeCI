package controllers

import (
	// "beerbubble/MtimeCI/datatype"
	"beerbubble/MtimeCI/models"
	"beerbubble/MtimeCI/utility"
	// "fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
	//"github.com/astaxie/beego/logs"
)

type RundeckController struct {
	beego.Controller
}

func (this *RundeckController) ServerList() {

	utility.ViewLogin(this.Ctx)

	o := orm.NewOrm()

	var servers []*models.Rundeckserver
	o.QueryTable("Rundeckserver").All(&servers)

	this.Data["Title"] = "Rundeck Server 列表"
	this.Data["servers"] = servers

	this.Layout = "Template.html"
	this.TplNames = "rundeck/list.html"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["NavContent"] = "component/nav.html"
	this.LayoutSections["Scripts"] = "env/listjs.html"
	this.LayoutSections["HtmlHead"] = "env/listcss.html"
}

func (this *RundeckController) AddOrEdit() {

	utility.ViewLogin(this.Ctx)

	serverid := this.Input().Get("serverid")

	o := orm.NewOrm()

	var server models.Rundeckserver

	var title string

	if serverid == "" {
		title = "添加Rundeck Server"

	} else {
		title = "编辑Rundeck Server"

		//var env models.Environmentinfo
		o.QueryTable("Rundeckserver").Filter("id", serverid).One(&server)
	}

	this.Data["Title"] = title
	this.Data["server"] = server

	this.Layout = "Template.html"
	this.TplNames = "rundeck/addoredit.html"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["NavContent"] = "component/nav.html"
	this.LayoutSections["Scripts"] = "rundeck/addoreditjs.html"
	this.LayoutSections["HtmlHead"] = "env/listcss.html"
}

func (this *RundeckController) AddApi() {

	serverid, _ := strconv.Atoi(this.Input().Get("serverid"))
	serverapiurl := this.Input().Get("serverapiurl")
	serverauthtoken := this.Input().Get("serverauthtoken")

	o := orm.NewOrm()

	if serverid > 0 {
		server := models.Rundeckserver{Id: serverid}
		if o.Read(&server) == nil {
			server.Apiurl = serverapiurl
			server.Authtoken = serverauthtoken

			if num, err := o.Update(&server); err == nil {
				this.Data["json"] = models.EnvAddModel{models.JsonResultBaseStruct{Result: true, Message: "操作成功"}, num}
			}
		}

	} else {
		var server models.Rundeckserver
		server.Apiurl = serverapiurl
		server.Authtoken = serverauthtoken

		if id, err := o.Insert(&server); err == nil {
			this.Data["json"] = models.EnvAddModel{models.JsonResultBaseStruct{Result: true, Message: "操作成功"}, id}
		}
	}

	this.ServeJson()
}

func (this *RundeckController) AddOrEditJob() {

	utility.ViewLogin(this.Ctx)

	rundeckjobid := this.Input().Get("rundeckjobid")
	moduleid := this.Input().Get("moduleid")

	o := orm.NewOrm()

	var job models.Rundeckjob

	var title string

	if rundeckjobid == "" {
		title = "添加Rundeck Job"

	} else {
		title = "编辑Rundeck Job"

		o.QueryTable("Rundeckjob").Filter("id", rundeckjobid).One(&job)
	}

	var module models.Moduleinfo
	o.QueryTable("Moduleinfo").Filter("id", moduleid).One(&module)

	var envs []*models.Environmentinfo
	o.QueryTable("Environmentinfo").All(&envs)

	viewEnvironmentinfoModels := []*models.ViewEnvironmentinfoModel{}
	for i := 0; i < len(envs); i++ {
		selected := ""

		if envs[i].Id == job.Envid {
			selected = "selected"
		}

		viewEnvironmentinfoModels = append(viewEnvironmentinfoModels, &models.ViewEnvironmentinfoModel{envs[i].Id, envs[i].Name, envs[i].Description, envs[i].Rundeckapiurl, selected})
	}

	var servers []*models.Rundeckserver
	o.QueryTable("Rundeckserver").All(&servers)

	viewRundeckServerModels := []*models.ViewRundeckServerModel{}
	for i := 0; i < len(servers); i++ {
		selected := ""

		if servers[i].Id == job.Rundeckserverid {
			selected = "selected"
		}

		viewRundeckServerModels = append(viewRundeckServerModels, &models.ViewRundeckServerModel{servers[i].Id, servers[i].Apiurl, selected})
	}

	this.Data["Title"] = title
	this.Data["job"] = job
	this.Data["module"] = module
	this.Data["envs"] = viewEnvironmentinfoModels
	this.Data["servers"] = viewRundeckServerModels

	this.Layout = "Template.html"
	this.TplNames = "rundeck/addoreditjob.html"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["NavContent"] = "component/nav.html"
	this.LayoutSections["Scripts"] = "rundeck/addoreditjobjs.html"
	this.LayoutSections["HtmlHead"] = "env/listcss.html"
}

func (this *RundeckController) AddJobApi() {

	moduleid, _ := strconv.Atoi(this.Input().Get("moduleid"))
	jobid, _ := strconv.Atoi(this.Input().Get("jobid"))
	envid, _ := strconv.Atoi(this.Input().Get("envid"))
	rundeckserverid, _ := strconv.Atoi(this.Input().Get("rundeckserverid"))

	rundeckbuildjobid := this.Input().Get("rundeckbuildjobid")
	rundeckpackagejobid := this.Input().Get("rundeckpackagejobid")

	o := orm.NewOrm()

	if jobid > 0 {
		job := models.Rundeckjob{Id: jobid}
		if o.Read(&job) == nil {
			job.Moduleid = moduleid
			job.Envid = envid
			job.Rundeckserverid = rundeckserverid
			job.Rundeckbuildjobid = rundeckbuildjobid
			job.Rundeckpackagejobid = rundeckpackagejobid

			if num, err := o.Update(&job); err == nil {
				this.Data["json"] = models.EnvAddModel{models.JsonResultBaseStruct{Result: true, Message: "操作成功"}, num}
			}
		}

	} else {
		var job models.Rundeckjob
		job.Moduleid = moduleid
		job.Envid = envid
		job.Rundeckserverid = rundeckserverid
		job.Rundeckbuildjobid = rundeckbuildjobid
		job.Rundeckpackagejobid = rundeckpackagejobid

		if id, err := o.Insert(&job); err == nil {
			this.Data["json"] = models.EnvAddModel{models.JsonResultBaseStruct{Result: true, Message: "操作成功"}, id}
		}
	}

	this.ServeJson()
}
