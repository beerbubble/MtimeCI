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
	"os/exec"
	"strings"
)

type ProjectController struct {
	beego.Controller
}

func (this *ProjectController) Get() {
	utility.ViewLogin(this.Ctx)

	this.Layout = "Template.html"
	this.TplNames = "project/index.html"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["NavContent"] = "component/nav.html"
}

func (this *ProjectController) List() {

	utility.ViewLogin(this.Ctx)
	languageType := this.Input().Get("languageType")

	o := orm.NewOrm()

	var projects []models.Projectinfo
	o.QueryTable("Projectinfo").Filter("LanguageType", languageType).All(&projects)
	//fmt.Printf("Returned Rows Num: %s, %s", num2, err2)
	//.Filter("LanguageType", languageType)

	var title string
	switch languageType {
	case "1":
		title = "Go Projects"
	case "2":
		title = "C# Projects"
	case "3":
		title = "Java Projects"
	}

	viewModel := []*models.ViewProjectListModel{}

	for i := 0; i < len(projects); i++ {
		var projectbranchs []*models.Projectbranch
		//num, _ :=
		o.QueryTable("Projectbranch").Filter("Projectid", projects[i].Id).All(&projectbranchs)
		//fmt.Printf("Branch Numbers : %s", num)
		viewModel = append(viewModel, &models.ViewProjectListModel{projects[i], projectbranchs})
	}

	this.Data["Title"] = title
	//this.Data["projectsnum"] = num
	this.Data["projects"] = projects
	this.Data["viewModel"] = viewModel

	this.Layout = "Template.html"
	this.TplNames = "project/list.html"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["NavContent"] = "component/nav.html"
	this.LayoutSections["Scripts"] = "project/listjs.html"
	this.LayoutSections["HtmlHead"] = "project/listcss.html"
}

func (this *ProjectController) UpdateBranch() {
	projectid := this.Input().Get("projectid")
	intprojectid, _ := strconv.Atoi(projectid)

	o := orm.NewOrm()

	var project models.Projectinfo
	o.QueryTable("Projectinfo").Filter("Id", projectid).One(&project)

	cmd := exec.Command("git", "branch", "-r")
	cmd.Dir = project.Sourceurl //"/Users/Liujia/Work/Mtime/Git/go/basis/config-web/"
	out, _ := cmd.Output()

	fmt.Printf("%s\n", out)

	branchs := strings.Split(strings.TrimSpace(strings.Replace(string(out), "origin/", "", -1)), "\n  ")

	fmt.Println(strings.Replace(string(out), "origin/", "", -1))

	branchs = append(branchs[:0], branchs[1:]...)

	//获取已有分支数据
	var projectbranchs []*models.Projectbranch
	num, _ := o.QueryTable("Projectbranch").Filter("Projectid", projectid).All(&projectbranchs)

	if num > 0 {
		for i := 0; i < len(branchs); i++ {

			isadd := true
			for j := 0; j < len(projectbranchs); j++ {
				if branchs[i] == projectbranchs[j].Branchname {
					isadd = false
				}
			}
			if isadd {
				var branch models.Projectbranch
				branch.Projectid = intprojectid
				branch.Branchname = branchs[i]
				if id, err := o.Insert(&branch); err == nil {
					fmt.Printf("%s\n", id)
				}
			}
		}

	} else {
		for i := 0; i < len(branchs); i++ {
			var branch models.Projectbranch
			branch.Projectid = intprojectid
			branch.Branchname = branchs[i]
			if id, err := o.Insert(&branch); err == nil {
				fmt.Printf("%s\n", id)
			}
		}

	}

	result := models.ApiUpdateBranchModel{models.JsonResultBaseStruct{Result: true, Message: "更新分支列表成功"}, branchs}

	this.Data["json"] = &result //&UserList
	this.ServeJson()
}

func (this *ProjectController) Detail() {

	projectid := this.Ctx.Input.Params["0"]

	o := orm.NewOrm()
	var project models.Projectinfo
	o.QueryTable("Projectinfo").Filter("Id", projectid).One(&project)

	var projectbranchs []*models.Projectbranch
	o.QueryTable("Projectbranch").Filter("Projectid", projectid).All(&projectbranchs)

	var projectenvs []*models.Projectenvironment
	o.QueryTable("Projectenvironment").Filter("Projectid", projectid).OrderBy("EnvId").All(&projectenvs)

	var envs []*models.Environmentinfo
	o.QueryTable("Environmentinfo").All(&envs)

	var modules []*models.Moduleinfo
	o.QueryTable("Moduleinfo").Filter("Projectid", projectid).All(&modules)

	envmap := make(map[int]*models.Environmentinfo)
	for i := 0; i < len(envs); i++ {
		envmap[envs[i].Id] = envs[i]
	}

	var users []*models.User
	o.QueryTable("User").All(&users)

	usermap := make(map[int]string)
	for i := 0; i < len(users); i++ {
		usermap[users[i].Id] = users[i].Username
	}

	viewlocalprojectenvmodels := []*models.ViewProjectEnvironmentModel{}
	viewpreprojectenvmodels := []*models.ViewProjectEnvironmentModel{}
	viewonlineprojectenvmodels := []*models.ViewProjectEnvironmentModel{}

	viewprojectenvmodels := []*models.ViewProjectEnvironmentModel{}
	for i := 0; i < len(projectenvs); i++ {
		env := envmap[projectenvs[i].Envid]

		switch env.Envtype {
		case 1:
			viewlocalprojectenvmodels = append(viewlocalprojectenvmodels, &models.ViewProjectEnvironmentModel{projectenvs[i].Id, projectenvs[i].Projectid, projectenvs[i].Envid, projectenvs[i].Rundeckbuildjobid, projectenvs[i].Rundeckpackagejobid, projectenvs[i].Lastexcutiontime.Format("2006-01-02 15:04:05"), projectenvs[i].Lastexcutionuserid, envmap[projectenvs[i].Envid].Name, usermap[projectenvs[i].Lastexcutionuserid], projectenvs[i].Lastbuildnumber, projectenvs[i].Lastbuildbranchname, projectenvs[i].Lastbuildbranchhash, modules})
		case 2:
			viewpreprojectenvmodels = append(viewpreprojectenvmodels, &models.ViewProjectEnvironmentModel{projectenvs[i].Id, projectenvs[i].Projectid, projectenvs[i].Envid, projectenvs[i].Rundeckbuildjobid, projectenvs[i].Rundeckpackagejobid, projectenvs[i].Lastexcutiontime.Format("2006-01-02 15:04:05"), projectenvs[i].Lastexcutionuserid, envmap[projectenvs[i].Envid].Name, usermap[projectenvs[i].Lastexcutionuserid], projectenvs[i].Lastbuildnumber, projectenvs[i].Lastbuildbranchname, projectenvs[i].Lastbuildbranchhash, modules})
		case 3:
			viewonlineprojectenvmodels = append(viewonlineprojectenvmodels, &models.ViewProjectEnvironmentModel{projectenvs[i].Id, projectenvs[i].Projectid, projectenvs[i].Envid, projectenvs[i].Rundeckbuildjobid, projectenvs[i].Rundeckpackagejobid, projectenvs[i].Lastexcutiontime.Format("2006-01-02 15:04:05"), projectenvs[i].Lastexcutionuserid, envmap[projectenvs[i].Envid].Name, usermap[projectenvs[i].Lastexcutionuserid], projectenvs[i].Lastbuildnumber, projectenvs[i].Lastbuildbranchname, projectenvs[i].Lastbuildbranchhash, modules})
		}

		viewprojectenvmodels = append(viewprojectenvmodels, &models.ViewProjectEnvironmentModel{projectenvs[i].Id, projectenvs[i].Projectid, projectenvs[i].Envid, projectenvs[i].Rundeckbuildjobid, projectenvs[i].Rundeckpackagejobid, projectenvs[i].Lastexcutiontime.Format("2006-01-02 15:04:05"), projectenvs[i].Lastexcutionuserid, envmap[projectenvs[i].Envid].Name, usermap[projectenvs[i].Lastexcutionuserid], projectenvs[i].Lastbuildnumber, projectenvs[i].Lastbuildbranchname, projectenvs[i].Lastbuildbranchhash, modules})
	}

	fmt.Println(len(viewlocalprojectenvmodels))
	fmt.Println(len(viewpreprojectenvmodels))
	fmt.Println(len(viewonlineprojectenvmodels))

	viewModel := models.ViewProjectListModel{project, projectbranchs}

	//language := datatype.LanguageType(1)

	this.Data["ProjectId"] = projectid
	this.Data["viewModel"] = viewModel
	this.Data["LanguageType"] = datatype.LanguageTypeMap[project.Languagetype]
	this.Data["localprojectenvs"] = viewlocalprojectenvmodels
	this.Data["preprojectenvs"] = viewpreprojectenvmodels
	this.Data["onlineprojectenvs"] = viewonlineprojectenvmodels
	this.Data["modules"] = modules

	this.Layout = "Template.html"
	this.TplNames = "project/detail.html"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["NavContent"] = "component/nav.html"
	//this.LayoutSections["Scripts"] = "project/listjs.html"
	this.LayoutSections["HtmlHead"] = "project/listcss.html"
}
