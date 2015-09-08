package models

//import "github.com/astaxie/beego/orm"
import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
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
	Lastexcutiontime    time.Time
	Lastexcutionuserid  int
	Lastbuildnumber     int
	Lastbuildbranchname string
	Lastbuildbranchhash string
}

type Projectbuild struct {
	Id          int
	Projectid   int
	Buildnumber int
	Branchname  string
	Branchhash  string
	Buildpath   string
	Created     time.Time
	Buildstatus int
	Executionid int
}

func init() {
	orm.RegisterModel(new(Executionlog))
	orm.RegisterModel(new(User))
	orm.RegisterModel(new(Projectinfo))
	orm.RegisterModel(new(Environmentinfo))
	orm.RegisterModel(new(Projectbranch))
	orm.RegisterModel(new(Projectenvironment))
	orm.RegisterModel(new(Projectbuild))
}
