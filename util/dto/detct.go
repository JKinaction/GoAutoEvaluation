package dto

type CodeDto struct {
	Userid    string
	Problemid string
	Code      string
}

type FuncVarDto struct {
	Userid    string
	Problemid string
	Funcs     []string
	Vars      []string
	Code      string
}
