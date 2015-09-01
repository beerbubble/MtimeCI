package models

//import "github.com/astaxie/beego/orm"
import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type Executionlog struct {
	Id              int
	Projectid       int
	Envid           int
	Packagepath     string
	Projectbuildid  int
	Executionuserid int
	Createtime      string
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
	Id                  int
	Name                string
	Description         string
	Rundeckapiurl       string
	Rundeckapiauthtoken string
	Envtype             int
}

type Projectbranch struct {
	Id         int
	Projectid  int
	Branchname string
}

type Projectenvironment struct {
	Id                  int
	Projectid           int
	Envid               int
	Rundeckbuildjobid   string
	Rundeckpackagejobid string
	Lastexcutiontime    string
	Lastexcutionuserid  int
	Lastbuildnumber     int
	Lastbuildbranchname string
	Lastbuildbranchhash string
}

func init() {
	orm.RegisterModel(new(Executionlog))
	orm.RegisterModel(new(User))
	orm.RegisterModel(new(Projectinfo))
	orm.RegisterModel(new(Environmentinfo))
	orm.RegisterModel(new(Projectbranch))
	orm.RegisterModel(new(Projectenvironment))
}
