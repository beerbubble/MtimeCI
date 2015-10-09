package controllers

import (
	// "beerbubble/MtimeCI/models"
	// "beerbubble/MtimeCI/utility"
	// "fmt"
	"github.com/astaxie/beego"
	// "github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"

	//"io/ioutil"
	//"net/http"
	// "encoding/xml"
	//"net/url"
	// "time"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {

	this.Data["Website"] = "beego.me" //+ string(body)
	this.Data["Email"] = "astaxie@gmail.com"
	this.TplNames = "index.html"
	this.Layout = "Template.html"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["NavContent"] = "component/nav.html"
}
