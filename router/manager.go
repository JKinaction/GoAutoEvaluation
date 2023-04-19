package router

import "github.com/gin-gonic/gin"

func GetManager(route *gin.RouterGroup) {
	manager := route.Group("/manager")
	{
		manager.POST("")
		manager.POST("")
	}
}
