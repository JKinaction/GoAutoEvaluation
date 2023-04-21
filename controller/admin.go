package controller

import (
	"domo1/service"
	"domo1/util/dto"
	"domo1/util/response"

	"github.com/gin-gonic/gin"
)

func QuestionList(ctx *gin.Context) {

	res := service.QuestionList()
	response.HandleResponse(ctx, res)
}

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
	res := service.InputAnswerPublish(request)
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
	res := service.InputAnswerList(request)
	response.HandleResponse(ctx, res)
}
func InputAnswerDelete(ctx *gin.Context) {
	var request dto.InputAnswerIdDto
	err := ctx.Bind(&request)
	if err != nil {
		response.Fail(ctx, nil, response.RequestError)
		return
	}
	if request.Id == 0 {
		response.Fail(ctx, nil, response.RequestError)
		return
	}
	res := service.InputAnswerDelete(request)
	response.HandleResponse(ctx, res)
}
