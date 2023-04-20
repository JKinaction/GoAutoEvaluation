package service

import (
	"domo1/util/common"
	"domo1/util/dto"
	"domo1/util/model"
	"domo1/util/response"
	"fmt"
	"net/http"
)

func InputAnswerPubblish(dto dto.InputAnswerDto) response.ResponseStruct {
	res := response.NewResponse()
	questionid := dto.Questionid
	ia := model.InputAnswer{Questionid: questionid}
	err := common.GetDB().Create(&ia).Error
	if err != nil {
		res.HttpStatus = http.StatusInternalServerError
		res.Code = response.ServerErrorCode
		res.Msg = response.SystemError
		return res
	}
	//保存到./file/question/Questionid/Intputanwserid
	//该路径下保存user.in和answer.out
	dir := fmt.Sprintf("./file/question/%v/%v/", questionid, ia.ID)

	return res
}

func QuestionGet(dto dto.QuestionDto) response.ResponseStruct {
	res := response.NewResponse()
	q := model.Question{}
	err := common.GetDB().Select(&q).Where("id=?", dto.Questionid).Error
	if err != nil {
		res.HttpStatus = http.StatusInternalServerError
		res.Code = response.ServerErrorCode
		res.Msg = response.SystemError
		return res
	}
	return res
}

func QuestionPublish(dto dto.QuestionDto) response.ResponseStruct {
	res := response.NewResponse()
	contents := dto.Contents
	q := model.Question{Contents: contents}
	err := common.GetDB().Create(&q).Error
	if err != nil {
		res.HttpStatus = http.StatusInternalServerError
		res.Code = response.ServerErrorCode
		res.Msg = response.SystemError
		return res
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
