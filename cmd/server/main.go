package main

import (
	"os"
	"simpro/config"
	"simpro/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	config.LoadEnv()

	r := gin.Default()

	routes.SetupRouter(r)

	port := os.Getenv("APP_PORT")

	if port == "" {
		port = "8080"
	}

	r.Run(":" + port)
}
