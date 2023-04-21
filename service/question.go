package service

import (
	"domo1/util/common"
	"domo1/util/dto"
	"domo1/util/model"
	"domo1/util/response"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func InputAnswerList(dto dto.InputAnswerDto) response.ResponseStruct {
	res := response.NewResponse()
	var ialist []model.InputAnswer
	err := common.GetDB().Where("questionid=?", dto.Questionid).Find(&ialist).Error
	if err != nil {
		logrus.Println(err)
		res.HttpStatus = http.StatusInternalServerError
		res.Code = response.ServerErrorCode
		res.Msg = response.SystemError
		return res
	}

	return res
}

func InputAnswerDelete(dto dto.InputAnswerIdDto) response.ResponseStruct {
	res := response.NewResponse()
	var ia model.InputAnswer
	err := common.GetDB().Delete(&ia).Where("id=?", dto.Id).Error
	if err != nil {
		logrus.Println(err)
		res.HttpStatus = http.StatusInternalServerError
		res.Code = response.ServerErrorCode
		res.Msg = response.SystemError
		return res
	}

	return res
}

func InputAnswerPublish(dto dto.InputAnswerDto) response.ResponseStruct {
	res := response.NewResponse()
	var q model.Question
	questionid := dto.Questionid
	common.GetDB().Where("id=?", questionid).First(&q)
	if q.ID == 0 {
		logrus.Println(q)
		res.HttpStatus = http.StatusBadRequest
		res.Code = response.FailCode
		res.Msg = response.ParameterError
		return res
	}
	ia := model.InputAnswer{Questionid: questionid}
	err := common.GetDB().Create(&ia).Error
	if err != nil {
		logrus.Println(err)
		res.HttpStatus = http.StatusInternalServerError
		res.Code = response.ServerErrorCode
		res.Msg = response.SystemError
		return res
	}
	//保存到./file/question/Questionid/Intputanwserid
	//该路径下保存user.in和answer.out
	dir := fmt.Sprintf("./file/question/%v/%v/", questionid, ia.ID)

	if err := os.MkdirAll(fmt.Sprintf("./file/question/%v/%v", questionid, ia.ID), os.ModePerm); err != nil {
		logrus.Println(err)
		res.HttpStatus = http.StatusInternalServerError
		res.Code = response.ServerErrorCode
		res.Msg = response.SystemError
		return res
	}
	infile, err := os.Create(dir + "user.in")
	if err != nil {
		logrus.Println(err)
		res.HttpStatus = http.StatusInternalServerError
		res.Code = response.ServerErrorCode
		res.Msg = response.SystemError
		return res
	}
	defer infile.Close()
	os.WriteFile(dir+"user.in", []byte(dto.Input), 0777)
	answerfile, err := os.Create(dir + "answer.out")
	if err != nil {
		logrus.Println(err)
		res.HttpStatus = http.StatusInternalServerError
		res.Code = response.ServerErrorCode
		res.Msg = response.SystemError
		return res
	}
	defer answerfile.Close()
	os.WriteFile(dir+"answer.out", []byte(dto.Answer), 0777)
	err = common.GetDB().Model(&model.InputAnswer{}).Where("id=?", ia.ID).Update("path", fmt.Sprintf("./file/question/%v/%v", questionid, ia.ID)).Error
	if err != nil {
		logrus.Println(err)
		res.HttpStatus = http.StatusInternalServerError
		res.Code = response.ServerErrorCode
		res.Msg = response.SystemError
		return res
	}
	return res
}

func QuestionGet(dto dto.QuestionDto) response.ResponseStruct {
	res := response.NewResponse()
	q := model.Question{}
	err := common.GetDB().Where("id=?", dto.Questionid).First(&q).Error
	if err != nil {
		res.HttpStatus = http.StatusInternalServerError
		res.Code = response.ServerErrorCode
		res.Msg = response.SystemError
		return res
	}
	res.Data = gin.H{
		"question": q,
	}
	return res
}

func QuestionPublish(dto dto.QuestionDto) response.ResponseStruct {
	res := response.NewResponse()
	contents := dto.Contents
	q := model.Question{Contents: contents, Name: dto.Name}
	err := common.GetDB().Create(&q).Error
	if err != nil {
		res.HttpStatus = http.StatusInternalServerError
		res.Code = response.ServerErrorCode
		res.Msg = response.SystemError
		return res
	}
	res.Data = gin.H{
		"question": q.ID,
	}
	return res
}
func QuestionUpdate(dto dto.QuestionDto) response.ResponseStruct {
	res := response.NewResponse()
	err := common.GetDB().Update("contents", dto.Contents).Where("id= ?", dto.Questionid).Error
	if err != nil {
		res.HttpStatus = http.StatusInternalServerError
		res.Code = response.ServerErrorCode
		res.Msg = response.SystemError
		return res
	}
	return res
}

func QuestionDelete(dto dto.QuestionDto) response.ResponseStruct {
	res := response.NewResponse()
	var question model.Question
	err := common.GetDB().Delete(&question).Where("id= ?", dto.Questionid).Error
	if err != nil {
		res.HttpStatus = http.StatusInternalServerError
		res.Code = response.ServerErrorCode
		res.Msg = response.SystemError
		return res
	}
	return res
}

func QuestionList() response.ResponseStruct {
	res := response.NewResponse()
	var questions []model.Question
	err := common.GetDB().Find(&questions).Error
	if err != nil {
		res.HttpStatus = http.StatusInternalServerError
		res.Code = response.ServerErrorCode
		res.Msg = response.SystemError
		return res
	}
	res.Data = gin.H{
		"question": questions,
	}
	return res
}
