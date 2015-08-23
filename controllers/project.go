package controllers

import (
	//"beerbubble/MtimeCI/datatype"
	"beerbubble/MtimeCI/models"
	"beerbubble/MtimeCI/utility"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	//"github.com/astaxie/beego/logs"
	//"strconv"
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
	this.Data["Title"] = title
	this.Data["projectsnum"] = num
	this.Data["projects"] = projects

	this.Layout = "Template.html"
	this.TplNames = "project/list.html"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["NavContent"] = "component/nav.html"
	this.LayoutSections["Scripts"] = "project/listjs.html"
	this.LayoutSections["HtmlHead"] = "project/listcss.html"
}

func (this *ProjectController) UpdateBranch() {
	projectid := this.Input().Get("projectid")

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

	this.Data["json"] = &branchs //&UserList
	this.ServeJson()
}
