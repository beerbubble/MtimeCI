package models

type ViewProjectListModel struct {
	Project Projectinfo
	Branchs []*Projectbranch
}

type ViewProjectEnvironmentModel struct {
	Id                   int
	Projectid            int
	Envid                int
	Rundeckjobid         string
	Lastexcutiontime     string
	Lastexcutionuserid   int
	EnvName              string
	Lastexcutionusername string
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
