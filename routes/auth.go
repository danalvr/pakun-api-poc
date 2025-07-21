package routes

import (
	"net/http"
	"pakun-api-poc/services"

	"github.com/gin-gonic/gin"
)

type OTPRequest struct {
	Identifier string `json:"identifier"`
}

type OTPVerifyRequest struct {
	Identifier string `json:"identifier"`
	Code       string `json:"code"`
}

func RequestOTP(c *gin.Context) {
	var req OTPRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	otp := services.GenerateAndSaveOTP(req.Identifier)
	c.JSON(http.StatusOK, gin.H{"message": "OTP sent", "otp": otp})
}

func VerifyOTP(c *gin.Context) {
	var req OTPVerifyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	if services.VerifyOTP(req.Identifier, req.Code) {
		c.JSON(http.StatusOK, gin.H{"message": "Login success"})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired OTP"})
	}
}