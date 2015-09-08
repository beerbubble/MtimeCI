package datatype

import (
//"strconv"
)

/*
type LanguageType int

const (
	Golang = 1 + iota
	CSharp
	Java
)

var LanguageTypes = [...]string{
	"Golang",
	"CSharp",
	"Java",
}

func (l LanguageType) String() string {
	return LanguageTypes[l-1]
}
*/

var LanguageTypeMap = map[int]string{
	1: "Golang",
	2: "CSharp",
	3: "Java",
}

var EnvTypeMap = map[int]string{
	1: "Local",
	2: "PreOnline",
	3: "Online",
}

var BuildMap = map[int]string{
	1: "Start",
	2: "Running",
	3: "Success",
	4: "Fail",
}
