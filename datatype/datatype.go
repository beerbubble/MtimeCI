package datatype

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

func (l LanguageType) String() string { return LanguageTypes[l-1] }
