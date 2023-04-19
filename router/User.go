package router

import (
	"demo1/controller"

	"github.com/gin-gonic/gin"
)

func GetUserRoutes(route *gin.RouterGroup) {
	user := route.Group("/user")
	{
		user.POST("/register", controller.Register)
		user.POST("/login", controller.Login)
	}

	question := route.Group("question")
	{
		question.GET("/question", controller.GetQuestion)
		question.GET("/question-list", controller.GetQuestionList)
	}
}
