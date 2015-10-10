package models

type EnvAddModel struct {
	JsonResultBaseStruct
	Data int64
}

type ViewEnvTypeModel struct {
	Id       int
	Name     string
	Selected string
}

type ViewEnvironmentinfoItemModel struct {
	Id                  int
	Name                string
	Description         string
	Rundeckapiurl       string
	Rundeckapiauthtoken string
	Envtype             int
	Envtypename         string
}
