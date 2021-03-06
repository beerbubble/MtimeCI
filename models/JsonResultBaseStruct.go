package models

type JsonResultBaseStruct struct {
	Result  bool
	Message string
}

type BuildApiModel struct {
	JsonResultBaseStruct
	Data string
}

type ExecutionStatusApiModel struct {
	JsonResultBaseStruct
	Data string
}
