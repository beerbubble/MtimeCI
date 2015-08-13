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

func init() {
	orm.RegisterModel(new(Executionlog))
	orm.RegisterModel(new(User))
}
