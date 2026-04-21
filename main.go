package main

import (
	"BlogPost/config"
	"BlogPost/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()

	r := gin.Default()
	routes.SetupRoutes(r, config.DB)

	r.Run(":8080")
}
