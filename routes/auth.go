package routes

import (
	"net/http"
	"pakun-api-poc/services"
	"pakun-api-poc/utils"

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
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	otp, err := services.GenerateAndSaveOTP(req.Identifier)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate OTP"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "OTP sent", "otp": otp})
}

func VerifyOTP(c *gin.Context) {
	var req OTPVerifyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	valid, _ := services.VerifyOTP(req.Identifier, req.Code)
	if !valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired OTP"})
		return
	}

	token, err := utils.GenerateJWT(req.Identifier)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login success", "token": token})
}