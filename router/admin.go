package router

import (
	"domo1/controller"

	"github.com/gin-gonic/gin"
)

func GetAdmin(route *gin.RouterGroup) {
	admin := route.Group("/admin")
	question := admin.Group("/question")
	{
		question.GET("/question-list", controller.QuestionList)
		question.POST("/get", controller.QuestionGet)
		question.POST("/publish", controller.QuestionPublish)
		question.POST("/update", controller.QuestionUpdate)
		question.POST("/delete", controller.QuestionDelete)
		ia := question.Group("/inputanswer")
		{
			ia.POST("/publish", controller.InputAnswerPublish)
			ia.GET("/list", controller.InputAnswerList)
			ia.POST("delete", controller.InputAnswerDelete)
		}
	}
}
