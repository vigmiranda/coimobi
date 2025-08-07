package main

import (
	"github.com/vigmiranda/coimobi-service/pkg/configs"
	"github.com/vigmiranda/coimobi-service/pkg/router"
)

func main() {
	configs.ConnectDatabase()
	r := router.SetupRoutes()
	r.Run(":8080")
}
