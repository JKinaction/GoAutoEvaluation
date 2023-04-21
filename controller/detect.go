package controller

import (
	"domo1/service"
	"domo1/util/dto"
	"domo1/util/response"

	"github.com/gin-gonic/gin"
)

func RunCode(c *gin.Context) {
	var request dto.CodeDto
	err := c.Bind(&request)
	if err != nil {
		response.Fail(c, nil, response.RequestError)
		return
	}
	res := service.RuncodeService(request)
	response.HandleResponse(c, res)
}

func CheckFuncVar(c *gin.Context) {
	var request dto.FuncVarDto
	err := c.Bind(&request)
	if err != nil {
		response.Fail(c, nil, response.RequestError)
		return
	}
	res := service.CheckFuncVarService(request)
	response.HandleResponse(c, res)
}
