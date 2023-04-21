package main

import (
	"domo1/router"
	"domo1/util/common"
)

func main() {
	common.InitDB()
	r := router.InitRouter()
	r.Run(":9000")
}
