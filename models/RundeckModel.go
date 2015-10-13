package models

type ViewRundeckServerModel struct {
	Id       int
	Apiurl   string
	Selected string
}

type ViewRundeckJobModel struct {
	Rundeckjobid        int
	Rundeckmoduleid     int
	Rundeckserverurl    string
	Rundeckbuildjobid   string
	Rundeckpackagejobid string
}
