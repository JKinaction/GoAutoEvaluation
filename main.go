package main

import (
	"demo1/util/logger"
	"domo1/router"
	"domo1/util/common"
)

func main() {
	logger.InitLogger()
	common.InitDB()
	r := router.InitRouter()
	r.Run(":9000")
}
