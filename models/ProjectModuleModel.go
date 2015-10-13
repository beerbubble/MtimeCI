package models

import (
	_ "time"
)

type ViewProjectModuleManageModel struct {
	Env Environmentinfo
	Job ViewRundeckJobModel
}

/*
type ViewProjectModuleJobModel struct {
	Name        string
	Description string
}
*/
