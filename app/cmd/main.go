package main

import (
	"dashboard/app/internal/router"
)

func main() {
	init := router.Init()
	init.Run(":8088")
}
