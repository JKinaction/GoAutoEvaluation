package router

import (
	"domo1/controller"

	"github.com/gin-gonic/gin"
)

func GetAdmin(route *gin.RouterGroup) {
	admin := route.Group("/admin")
	question := admin.Group("/question")
	{
		question.POST("/get", controller.QuestionGet)
		question.POST("/publish", controller.QuestionPublish)
		question.POST("update", controller.QuestionUpdate)
		question.POST("delete", controller.QuestionDelete)
	}
	ia := admin.Group("/inputanswer")
	{
		ia.POST("/publish", controller.InputAnswerPublish)
		ia.GET("/list")
	}
}
