package router

import (
	"domo1/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.CORSMiddleware())
	v1 := r.Group("/api")
	{
		GetUserRoutes(v1)
		GetDetect(v1)
		GetAdmin(v1)
	}
	return r
}
