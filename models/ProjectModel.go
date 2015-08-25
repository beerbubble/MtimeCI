package models

type ViewProjectListModel struct {
	Project *Projectinfo
	Branchs []*Projectbranch
}

type ApiUpdateBranchModel struct {
	JsonResultBaseStruct
	Data []string
}
