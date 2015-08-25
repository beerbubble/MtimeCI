package controllers

import (
	//"beerbubble/MtimeCI/datatype"
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

	var projects []*models.Projectinfo
	num, _ := o.QueryTable("Projectinfo").Filter("LanguageType", languageType).All(&projects)
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
		num, _ := o.QueryTable("Projectbranch").Filter("Projectid", projects[i].Id).All(&projectbranchs)
		fmt.Printf("Branch Numbers : %s", num)
		viewModel = append(viewModel, &models.ViewProjectListModel{projects[i], projectbranchs})
	}

	fmt.Printf("Number:%s", num)

	this.Data["Title"] = title
	this.Data["projectsnum"] = num
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
	err := o.QueryTable("Projectinfo").Filter("Id", projectid).One(&project)
	if err == orm.ErrMultiRows || err == orm.ErrNoRows {
	}

	cmd := exec.Command("git", "branch", "-r")
	cmd.Dir = project.Sourceurl //"/Users/Liujia/Work/Mtime/Git/go/basis/config-web/"
	out, _ := cmd.Output()

	fmt.Printf("%s\n", out)

	branchs := strings.Split(strings.TrimSpace(string(out)), "\n  ")

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

	projectid := this.Input().Get("projectid")

	this.Data["ProjectId"] = projectid
	this.Layout = "Template.html"
	this.TplNames = "project/detail.html"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["NavContent"] = "component/nav.html"
	this.LayoutSections["Scripts"] = "project/listjs.html"
	this.LayoutSections["HtmlHead"] = "project/listcss.html"
}
