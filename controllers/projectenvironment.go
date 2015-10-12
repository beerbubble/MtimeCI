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
	"encoding/xml"
	"io/ioutil"
	"time"
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
	this.LayoutSections["Scripts"] = "projectenv/buildjs.html"
	this.LayoutSections["HtmlHead"] = "env/listcss.html"
}

func (this *ProjectEnvironmentController) PublishPre() {

	utility.ViewLogin(this.Ctx)

	projectid := this.Input().Get("projectid")
	envid := this.Input().Get("envid")

	var projectenv models.Projectenvironment
	var env *models.Environmentinfo
	var modules []*models.Moduleinfo
	var title string

	o := orm.NewOrm()

	var envs []*models.Environmentinfo
	o.QueryTable("Environmentinfo").All(&envs)

	o.QueryTable("Projectenvironment").Filter("envid", envid).Filter("projectid", projectid).One(&projectenv)

	o.QueryTable("Moduleinfo").Filter("projectid", projectid).All(&modules)

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

	var projectbuilds []*models.Projectbuild
	o.QueryTable("Projectbuild").Filter("Projectid", projectid).Filter("Buildstatus", 3).OrderBy("-id").Limit(10).All(&projectbuilds)

	this.Data["Title"] = title
	this.Data["projectenv"] = projectenv
	this.Data["project"] = project
	this.Data["envs"] = viewenvironmentinfomodels
	this.Data["envid"] = envid
	this.Data["projectbranchs"] = projectbranchs
	this.Data["projectbuilds"] = projectbuilds
	this.Data["modules"] = modules

	this.Layout = "Template.html"
	this.TplNames = "projectenv/publishpre.html"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["NavContent"] = "component/nav.html"
	this.LayoutSections["Scripts"] = "projectenv/publishprejs.html"
	this.LayoutSections["HtmlHead"] = "env/listcss.html"
}

func (this *ProjectEnvironmentController) BuildApi() {

	//projectenvid := this.Input().Get("projectenvid")
	envid := this.Input().Get("envid")
	intenvid, _ := strconv.Atoi(envid)

	rundeckbuildjobid := this.Input().Get("rundeckbuildjobid")
	//rundeckpackagejobid := this.Input().Get("rundeckpackagejobid")
	projectid := this.Input().Get("projectid")
	branchname := this.Input().Get("branchname")

	o := orm.NewOrm()

	//添加版本号数据以及更新项目版本号
	var project models.Projectinfo

	o.QueryTable("Projectinfo").Filter("Id", projectid).Update(orm.Params{
		"BuildNumber": orm.ColValue(orm.Col_Add, 1),
	})

	o.QueryTable("Projectinfo").Filter("Id", projectid).One(&project)

	now := time.Now()

	fmt.Println(now)

	var pb models.Projectbuild
	pb.Projectid = project.Id
	pb.Buildnumber = project.Buildnumber
	pb.Branchname = branchname
	pb.Created = now
	pb.Buildstatus = 1
	pb.Envid = intenvid
	id, _ := o.Insert(&pb)

	//获取项目环境信息进行编译部署
	var env models.Environmentinfo
	o.QueryTable("Environmentinfo").Filter("id", envid).One(&env)

	args := map[string]string{"BUILD_NUMBER": strconv.Itoa(project.Buildnumber), "Branch_NAME": branchname, "Repository_Path": project.Repositorypath}

	response := utility.RundeckRunJob(env.Rundeckapiurl, env.Rundeckapiauthtoken, rundeckbuildjobid, args)
	fmt.Println(response)

	//parse rundeck api xml response
	r := models.RunJobExecutions{}
	xml_err := xml.Unmarshal([]byte(response), &r)

	if xml_err != nil {
		fmt.Printf("error: %v", xml_err)
		this.Data["json"] = xml_err //url.QueryEscape(logs[0].Packagepath)
		this.ServeJson()
	}

	if r.Exs != nil {
		o.QueryTable("Projectbuild").Filter("Id", id).Update(orm.Params{
			"BuildStatus": 2,
			"ExecutionId": r.Exs[0].Id,
		})
		fmt.Println(r.Exs[0].Id)
	}
	//rundeck执行完成之后读取git hash
	dat, _ := ioutil.ReadFile("/Volumes/ftproot/mtime/upversion/MtimeGoConfigWeb/2/config-web/GitBranchHash")
	//check(err)
	fmt.Print(string(dat))

	o.QueryTable("Projectbuild").Filter("Id", id).Update(orm.Params{
		"BranchHash": string(dat),
	})

	//fmt.Println(response)

	//fmt.Println(projectenvid)
	//fmt.Println(envid)
	//fmt.Println(rundeckbuildjobid)
	//fmt.Println(rundeckpackagejobid)
	//fmt.Println(projectid)
	//fmt.Println(branchname)

	this.Data["json"] = models.BuildApiModel{models.JsonResultBaseStruct{Result: true, Message: "操作成功"}, r.Exs[0].Id}
	this.ServeJson()
}

func (this *ProjectEnvironmentController) PublishPreApi() {

	//projectenvid := this.Input().Get("projectenvid")
	envid := this.Input().Get("envid")
	rundeckbuildjobid := this.Input().Get("rundeckbuildjobid")
	//	rundeckpackagejobid := this.Input().Get("rundeckpackagejobid")
	projectid := this.Input().Get("projectid")
	buildnumber := this.Input().Get("buildnumber")

	o := orm.NewOrm()

	var project models.Projectinfo
	o.QueryTable("Projectinfo").Filter("Id", projectid).One(&project)

	var preenv models.Environmentinfo
	o.QueryTable("Environmentinfo").Filter("id", envid).One(&preenv)

	var projectbuild models.Projectbuild
	o.QueryTable("Projectbuild").Filter("projectid", projectid).Filter("BuildNumber", buildnumber).One(&projectbuild)

	var localenv models.Environmentinfo
	o.QueryTable("Environmentinfo").Filter("id", projectbuild.Envid).One(&localenv)

	var localprojectenv models.Projectenvironment
	o.QueryTable("Projectenvironment").Filter("Projectid", projectid).Filter("Envid", projectbuild.Envid).One(&localprojectenv)

	//打包操作
	args := map[string]string{"BUILD_NUMBER": buildnumber, "Repository_Path": project.Repositorypath}

	response := utility.RundeckRunJob(localenv.Rundeckapiurl, localenv.Rundeckapiauthtoken, localprojectenv.Rundeckpackagejobid, args)

	fmt.Println(response)

	//parse rundeck api xml response
	r := models.RunJobExecutions{}
	xml_err := xml.Unmarshal([]byte(response), &r)

	if xml_err != nil {
		fmt.Printf("error: %v", xml_err)
		this.Data["json"] = xml_err //url.QueryEscape(logs[0].Packagepath)
		this.ServeJson()
	}

	//status_r := models.RunJobExecutions{}

	for i := 0; i < 1; {

		status_r := models.RunJobExecutions{}

		status_response := utility.RundeckExecutionInfo(localenv.Rundeckapiurl, localenv.Rundeckapiauthtoken, r.Exs[0].Id)

		status_xml_err := xml.Unmarshal([]byte(status_response), &status_r)

		if status_xml_err != nil {
			fmt.Printf("error: %v", status_xml_err)
			this.Data["json"] = status_xml_err
			this.ServeJson()
		}

		if status_r.Exs[0].Status == "succeeded" {
			i = 1
		} else {
			fmt.Println(status_r.Exs[0].Status)
		}

		time.Sleep(time.Second * 3)
	}

	//调用预上线环境rundeck
	pre_args := map[string]string{"url": project.Repositoryurl + buildnumber + "/pre/" + project.Packagefilename}

	prejob_response := utility.RundeckRunJob(preenv.Rundeckapiurl, preenv.Rundeckapiauthtoken, rundeckbuildjobid, pre_args)

	fmt.Println(prejob_response)

	//parse rundeck api xml response
	pre_r := models.RunJobExecutions{}
	pre_xml_err := xml.Unmarshal([]byte(prejob_response), &pre_r)

	if pre_xml_err != nil {
		fmt.Printf("error: %v", pre_xml_err)
		this.Data["json"] = pre_xml_err //url.QueryEscape(logs[0].Packagepath)
		this.ServeJson()
	}

	/*
		fmt.Println(projectenvid)
		fmt.Println(envid)
		fmt.Println(rundeckbuildjobid)
		fmt.Println(rundeckpackagejobid)
		fmt.Println(projectid)
		fmt.Println(buildnumber)
	*/

	this.Data["json"] = models.BuildApiModel{models.JsonResultBaseStruct{Result: true, Message: "操作成功"}, pre_r.Exs[0].Id} //status_r.Exs[0].Id}
	this.ServeJson()
}

func (this *ProjectEnvironmentController) ExecutionStatus() {

	executionid := this.Input().Get("executionid")
	envid := this.Input().Get("envid")
	projectenvid := this.Input().Get("projectenvid")

	o := orm.NewOrm()

	now := time.Now()

	fmt.Println(now)

	//获取项目环境信息进行编译部署
	var env models.Environmentinfo
	o.QueryTable("Environmentinfo").Filter("id", envid).One(&env)

	r := models.RunJobExecutions{}
	response := utility.RundeckExecutionInfo(env.Rundeckapiurl, env.Rundeckapiauthtoken, executionid)

	xml_err := xml.Unmarshal([]byte(response), &r)

	if xml_err != nil {
		fmt.Printf("error: %v", xml_err)
		this.Data["json"] = xml_err
		this.ServeJson()
	}

	switch r.Exs[0].Status {
	case "succeeded":
		o.QueryTable("Projectbuild").Filter("Executionid", executionid).Update(orm.Params{
			"BuildStatus": 3,
		})

		//获取项目环境信息进行编译部署
		var pb models.Projectbuild
		o.QueryTable("Projectbuild").Filter("Executionid", executionid).One(&pb)

		o.QueryTable("Projectenvironment").Filter("Id", projectenvid).Update(orm.Params{
			"Lastexcutiontime":    time.Now(),
			"LastBuildNumber":     pb.Buildnumber,
			"Lastbuildbranchname": pb.Branchname,
			"LastBuildBranchHash": pb.Branchhash,
		})
	case "failed":
		o.QueryTable("Projectbuild").Filter("Executionid", executionid).Update(orm.Params{
			"BuildStatus": 4,
		})
	}

	this.Data["json"] = models.ExecutionStatusApiModel{models.JsonResultBaseStruct{Result: true, Message: "操作成功"}, r.Exs[0].Status} //models.JsonResultBaseStruct{Result: true, Message: "操作成功"}
	this.ServeJson()
}
