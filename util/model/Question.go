package model

import "gorm.io/gorm"

type Question struct {
	gorm.Model
	Contents string
}

type InputAnswer struct {
	gorm.Model
	Path       string
	Questionid int
}
