package dto

type InputAnswerDto struct {
	Questionid int
	Input      string
	Answer     string
}

type QuestionDto struct {
	Questionid int
	Contents   string
}
