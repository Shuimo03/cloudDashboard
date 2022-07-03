package main

import (
	"dashboard/src/app/internal/router"
)

func main() {
	init := router.Init()
	init.Run(":8088")
}
