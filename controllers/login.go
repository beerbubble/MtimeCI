package controllers

import (
	"beerbubble/MtimeCI/models"
	"fmt"
	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Get() {
	this.Layout = "Template.html"
	this.TplNames = "login/login.html"

	this.LayoutSections = make(map[string]string)
	this.LayoutSections["Scripts"] = "login/loginjs.html"
}

func (this *LoginController) UserLogin() {
	email := this.Input().Get("email")
	password := this.Input().Get("password")

	o := orm.NewOrm()

	var user models.User
	err := o.QueryTable("User").Filter("Email", email).Filter("Password", password).One(&user)
	//fmt.Printf("Returned Rows Num: %s, %s", num, err)
	if err == orm.ErrMultiRows || err == orm.ErrNoRows {
		// 多条的时候报错
		fmt.Printf("Returned Multi Rows Not One")

		result := models.UserLoginModel{models.JsonResultBaseStruct{Result: false, Message: "账户名或密码错误!"}, user}

		this.Data["json"] = &result //&UserList
		this.ServeJson()
	}

	this.Ctx.SetCookie("MtimeCIUserId", strconv.Itoa(user.Id), 180*60, "/")

	result := models.UserLoginModel{models.JsonResultBaseStruct{Result: true, Message: "登录成功"}, user}

	this.Data["json"] = &result //&UserList
	this.ServeJson()
}
