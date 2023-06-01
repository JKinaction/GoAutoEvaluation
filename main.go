package main

import (
	"domo1/config"
	"domo1/router"
	"domo1/util/common"
	"domo1/util/logger"
)

func main() {
	config.InitConfig()
	logger.InitLogger()
	common.InitDB()
	r := router.InitRouter()
	r.Run(":9000")
}
