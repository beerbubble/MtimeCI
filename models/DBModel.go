package models

//import "github.com/astaxie/beego/orm"
import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type Executionlog struct {
	Id          int
	Projectname string
	Packagepath string
	Createtime  string
}

type User struct {
	Id         int
	Username   string
	Password   string
	Email      string
	Createtime string
}

type Projectinfo struct {
	Id           int
	Name         string
	Sourceurl    string
	Languagetype int
	Createtime   string
	Buildnumber  int
	Description  string
}

type Environmentinfo struct {
	Id            int
	Name          string
	Description   string
	Rundeckapiurl string
}

type Projectbranch struct {
	Id         int
	Projectid  int
	Branchname string
}

type ProjectEnvironment struct {
	Id        int
	Projectid int
	Envid     string
}

func init() {
	orm.RegisterModel(new(Executionlog))
	orm.RegisterModel(new(User))
	orm.RegisterModel(new(Projectinfo))
	orm.RegisterModel(new(Environmentinfo))
	orm.RegisterModel(new(Projectbranch))
	orm.RegisterModel(new(ProjectEnvironment))
}
