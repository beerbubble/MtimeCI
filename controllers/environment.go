package controllers

import (
	"beerbubble/MtimeCI/datatype"
	"beerbubble/MtimeCI/models"
	"beerbubble/MtimeCI/utility"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
	//"github.com/astaxie/beego/logs"
)

type EnvironmentController struct {
	beego.Controller
}

func (this *EnvironmentController) List() {

	utility.ViewLogin(this.Ctx)

	o := orm.NewOrm()

	var envs []*models.Environmentinfo
	o.QueryTable("Environmentinfo").All(&envs)

	// var envTypeList []string

	var viewEnvironmentinfoItemModel []*models.ViewEnvironmentinfoItemModel

	for i := 0; i < len(envs); i++ {
		env := envs[i]
		fmt.Println(datatype.EnvTypeMap[env.Envtype])
		// envTypeList = append(envTypeList, datatype.EnvTypeMap[env.Envtype])
		viewEnvironmentinfoItemModel = append(viewEnvironmentinfoItemModel, &models.ViewEnvironmentinfoItemModel{env.Id, env.Name, env.Description, env.Rundeckapiurl, env.Rundeckapiauthtoken, env.Envtype, datatype.EnvTypeMap[env.Envtype]})

	}

	// for key, value := range datatype.EnvTypeMap {
	// }

	this.Data["Title"] = "环境列表"
	//this.Data["envsnum"] = num
	// this.Data["envTypeList"] = envTypeList
	this.Data["envs"] = viewEnvironmentinfoItemModel

	this.Layout = "Template.html"
	this.TplNames = "env/list.html"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["NavContent"] = "component/nav.html"
	this.LayoutSections["Scripts"] = "env/listjs.html"
	this.LayoutSections["HtmlHead"] = "env/listcss.html"
}

func (this *EnvironmentController) AddOrEdit() {

	utility.ViewLogin(this.Ctx)

	envid := this.Input().Get("envid")

	var env models.Environmentinfo
	var title string

	//envTypes := datatype.EnvTypes

	if envid == "" {
		title = "添加平台信息"

	} else {
		title = "编辑平台信息"

		o := orm.NewOrm()

		//var env models.Environmentinfo
		o.QueryTable("Environmentinfo").Filter("id", envid).One(&env)
		//if err == orm.ErrMultiRows {
		// 多条的时候报错
		//fmt.Printf("Returned Multi Rows Not One")
		//}
	}

	viewEnvTypeModels := []*models.ViewEnvTypeModel{}

	for key, value := range datatype.EnvTypeMap {
		selected := ""
		//fmt.Println(env.Envtype)
		if key == env.Envtype {
			selected = "selected"
		}
		//fmt.Println(key, value, selected)
		viewEnvTypeModels = append(viewEnvTypeModels, &models.ViewEnvTypeModel{key, value, selected})
	}

	this.Data["Title"] = title
	this.Data["env"] = env
	this.Data["envTypes"] = viewEnvTypeModels

	this.Layout = "Template.html"
	this.TplNames = "env/addoredit.html"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["NavContent"] = "component/nav.html"
	this.LayoutSections["Scripts"] = "env/addoreditjs.html"
	this.LayoutSections["HtmlHead"] = "env/listcss.html"
}

func (this *EnvironmentController) AddApi() {

	envid := this.Input().Get("envid")
	envname := this.Input().Get("envname")
	envdes := this.Input().Get("envdes")
	envapiurl := this.Input().Get("envapiurl")
	envapiauthtoken := this.Input().Get("envapiauthtoken")
	envtype, _ := strconv.Atoi(this.Input().Get("envtype"))

	fmt.Println(envname)

	intenvid, _ := strconv.Atoi(envid)

	o := orm.NewOrm()

	if intenvid > 0 {
		env := models.Environmentinfo{Id: intenvid}
		if o.Read(&env) == nil {
			env.Name = envname
			env.Description = envdes
			env.Rundeckapiurl = envapiurl
			env.Rundeckapiauthtoken = envapiauthtoken
			env.Envtype = envtype
			if num, err := o.Update(&env); err == nil {
				this.Data["json"] = models.EnvAddModel{models.JsonResultBaseStruct{Result: true, Message: "操作成功"}, num}
			}
		}

	} else {

		var env models.Environmentinfo
		env.Name = envname
		env.Description = envdes
		env.Rundeckapiurl = envapiurl
		env.Rundeckapiauthtoken = envapiauthtoken
		env.Envtype = envtype
		if id, err := o.Insert(&env); err == nil {
			this.Data["json"] = models.EnvAddModel{models.JsonResultBaseStruct{Result: true, Message: "操作成功"}, id}
		}
	}

	this.ServeJson()
}
