package controller

import (
	"domo1/service"
	"domo1/util/dto"
	"domo1/util/response"

	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context) {
	var request dto.LoginDto
	err := ctx.Bind(&request)
	if err != nil {
		response.Fail(ctx, nil, response.RequestError)
		return
	}
	if len(request.Password) < 3 {
		response.CheckFail(ctx, nil, response.PasswordCheck)
		return
	}
	res := service.LoginService(request)
	response.HandleResponse(ctx, res)

}

func Register(ctx *gin.Context) {
	var request dto.RegisterDto
	err := ctx.Bind(&request)
	if err != nil {
		response.Fail(ctx, nil, response.RequestError)
		return
	}

	if len(request.Password) < 3 {
		response.CheckFail(ctx, nil, response.PasswordCheck)
		return
	}

	res := service.RegisterService(request)
	response.HandleResponse(ctx, res)
}
