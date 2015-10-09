package models

import (
	_ "time"
)

type ViewProjectListModel struct {
	Project Projectinfo
	Branchs []*Projectbranch
}

type ViewProjectEnvironmentModel struct {
	Id                   int
	Projectid            int
	Envid                int
	Rundeckbuildjobid    string
	Rundeckpackagejobid  string
	Lastexcutiontime     string
	Lastexcutionuserid   int
	EnvName              string
	Lastexcutionusername string
	Lastbuildnumber      int
	Lastbuildbranchname  string
	Lastbuildbranchhash  string
}

type ApiUpdateBranchModel struct {
	JsonResultBaseStruct
	Data []string
}

type ViewEnvironmentinfoModel struct {
	Id            int
	Name          string
	Description   string
	Rundeckapiurl string
	Selected      string
}
