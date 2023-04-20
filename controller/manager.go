package controller

import (
	"domo1/service"
	"domo1/util/dto"
	"domo1/util/response"

	"github.com/gin-gonic/gin"
)

func QuestionGet(ctx *gin.Context) {
	var request dto.QuestionDto
	err := ctx.Bind(&request)
	if err != nil {
		response.Fail(ctx, nil, response.RequestError)
		return
	}
	if request.Questionid == 0 {
		response.Fail(ctx, nil, response.RequestError)
		return
	}
	res := service.QuestionGet(request)
	response.HandleResponse(ctx, res)
}

func QuestionPublish(ctx *gin.Context) {
	var request dto.QuestionDto
	err := ctx.Bind(&request)
	if err != nil {
		response.Fail(ctx, nil, response.RequestError)
		return
	}
	if request.Contents == "" {
		response.Fail(ctx, nil, response.RequestError)
		return
	}
	res := service.QuestionPublish(request)
	response.HandleResponse(ctx, res)
}

func GetQuestionID(ctx *gin.Context) {
	var request dto.QuestionDto
	err := ctx.Bind(&request)
	if err != nil {
		response.Fail(ctx, nil, response.RequestError)
		return
	}
	if request.Questionid == 0 {
		response.Fail(ctx, nil, response.RequestError)
		return
	}
	res := service.QuestionGet(request)
	response.HandleResponse(ctx, res)
}

func QuestionUpdate(ctx *gin.Context) {
	var request dto.QuestionDto
	err := ctx.Bind(&request)
	if err != nil {
		response.Fail(ctx, nil, response.RequestError)
		return
	}
	if request.Questionid == 0 {
		response.Fail(ctx, nil, response.RequestError)
		return
	}
	res := service.QuestionUpdate(request)
	response.HandleResponse(ctx, res)
}

func QuestionDelete(ctx *gin.Context) {
	var request dto.QuestionDto
	err := ctx.Bind(&request)
	if err != nil {
		response.Fail(ctx, nil, response.RequestError)
		return
	}
	if request.Questionid == 0 {
		response.Fail(ctx, nil, response.RequestError)
		return
	}
	res := service.QuestionDelete(request)
	response.HandleResponse(ctx, res)
}

func InputAnswerPublish(ctx *gin.Context) {
	var request dto.InputAnswerDto
	err := ctx.Bind(&request)
	if err != nil {
		response.Fail(ctx, nil, response.RequestError)
		return
	}
	if request.Questionid == 0 {
		response.Fail(ctx, nil, response.RequestError)
		return
	}
	res := service.InputAnswerPubblish(request)
	response.HandleResponse(ctx, res)
}

func InputAnswerList(ctx *gin.Context) {
	var request dto.InputAnswerDto
	err := ctx.Bind(&request)
	if err != nil {
		response.Fail(ctx, nil, response.RequestError)
		return
	}
	if request.Questionid == 0 {
		response.Fail(ctx, nil, response.RequestError)
		return
	}

}
