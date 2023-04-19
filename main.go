package main

import (
	"domo1/router"
)

func mian() {
	r := router.InitRouter()
	r.Run(":8080")
}
