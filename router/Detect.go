package router

import (
	"domo1/controller"

	"github.com/gin-gonic/gin"
)

func GetDetect(route *gin.RouterGroup) {
	detect := route.Group("/detect")
	{
		detect.POST("/runcode", controller.RunCode)
		detect.POST("")
	}
}
