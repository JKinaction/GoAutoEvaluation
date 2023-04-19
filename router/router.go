package router

import "github.com/gin-gonic/gin"

func InitRouter() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/api/v1")
	{
		GetUserRoutes(v1)
		GetDetect(v1)
	}
	return r
}
