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
