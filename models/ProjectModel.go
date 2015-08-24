package models

type ViewProjectListModel struct {
	Id          int
	Name        string
	Sourceurl   string
	Buildnumber int
	Branchs     []*Projectbranch
}

type ApiUpdateBranchModel struct {
	JsonResultBaseStruct
	Data []string
}
