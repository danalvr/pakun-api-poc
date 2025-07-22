package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	server.POST("/auth/request-otp", RequestOTP)
	server.POST("/auth/verify-otp", VerifyOTP)
	server.GET("/dashboard", GetDashboard)
}