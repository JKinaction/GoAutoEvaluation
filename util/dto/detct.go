package dto

type CodeDto struct {
	Userid     int
	Questionid int
	Code       string
}

type FuncVarDto struct {
	Userid     int
	Questionid int
	Funcs      string
	Vars       string
	Signal     string
	Code       string
}
