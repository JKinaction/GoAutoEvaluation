package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Response(ctx *gin.Context, httpStatus int, code int, data gin.H, msg string) {
	ctx.JSON(httpStatus, gin.H{"code": code, "data": data, "msg": msg})
}

func Success(ctx *gin.Context, data gin.H, msg string) {
	Response(ctx, http.StatusOK, 2000, data, msg)
}

func Fail(ctx *gin.Context, data gin.H, msg string) {
	Response(ctx, http.StatusBadRequest, 4000, data, msg)
}

func CheckFail(ctx *gin.Context, data gin.H, msg string) {
	Response(ctx, http.StatusUnprocessableEntity, 4022, data, msg)
}

func HandleResponse(ctx *gin.Context, res ResponseStruct) {
	ctx.JSON(res.HttpStatus, gin.H{"code": res.Code, "data": res.Data, "msg": res.Msg})
}
func NewResponse() ResponseStruct {
	res := ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       SuccessCode,
		Data:       nil,
		Msg:        OK,
	}
	return res
}
