package services

import (
	"fmt"
	"math/rand"
	"pakun-api-poc/models"
	"time"
)

func generateRandomOTP() string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%06d", rand.Intn(1000000))
}

func GenerateAndSaveOTP(identifier string) string {
	otp := generateRandomOTP()
	models.OTPStore[identifier] = models.OTPEntry{
		Code: otp,
		ExpiresAt: time.Now().Add(time.Minute * 5),
	}

	return otp
}

func VerifyOTP(identifier string, code string) bool {
	entry, exists := models.OTPStore[identifier]
	if !exists {
		return false
	}
	if time.Now().After(entry.ExpiresAt) {
		delete(models.OTPStore, identifier)
		return false
	}
	if entry.Code != code {
		return false
	}
	delete(models.OTPStore, identifier)
	return true
}

