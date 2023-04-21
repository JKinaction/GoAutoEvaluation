package dto

type InputAnswerDto struct {
	Questionid int
	Input      string
	Answer     string
}

type QuestionDto struct {
	Questionid int
	Name       string
	Contents   string
}

type InputAnswerIdDto struct {
	Id int
}
