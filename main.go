package main

import (
	"pakun-api-poc/firebase"
	"pakun-api-poc/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	firebase.InitFirebase()
	defer firebase.Client.Close()

	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8000")
}