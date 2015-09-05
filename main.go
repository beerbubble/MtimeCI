package main

import (
	_ "beerbubble/MtimeCI/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterDriver("mysql", orm.DR_MySQL)

	orm.RegisterDataBase("default", "mysql", "root:root@tcp(localhost:8889)/MtimeCI?charset=utf8&loc=Asia%2FShanghai")

	orm.Debug = true
}

func main() {

	beego.Run()
}
