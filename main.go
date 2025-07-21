package main

import (
	"pakun-api-poc/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8000")
}