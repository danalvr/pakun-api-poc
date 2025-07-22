package services

import (
	"context"
	"fmt"
	"math/rand"
	"pakun-api-poc/firebase"
	"time"
)

type OTPEntry struct {
	Code string `firestore:"code"`
	ExpiresAt time.Time `firestore:"expiresAt"`
}

func generateRandomOTP() string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%06d", rand.Intn(1000000))
}

func GenerateAndSaveOTP(identifier string) (string, error) {
	otp := generateRandomOTP()
	entry := OTPEntry{
		Code: otp,
		ExpiresAt: time.Now().Add(time.Minute * 1),
	}

	_, err := firebase.Client.Collection("otps").Doc(identifier).Set(context.Background(), entry)
	if err != nil {
		return "", err
	}

	return otp, nil
}

func VerifyOTP(identifier string, code string) (bool, error) {
	doc, err := firebase.Client.Collection("otps").Doc(identifier).Get(context.Background())
	if err != nil {
		return false, nil
	}
	var entry OTPEntry
	doc.DataTo(&entry)

	if time.Now().After(entry.ExpiresAt) {
		firebase.Client.Collection("otps").Doc(identifier).Delete(context.Background())
		return false, nil
	}
	if entry.Code != code {
		return false, nil
	}
	firebase.Client.Collection("otps").Doc(identifier).Delete(context.Background())
	return true, nil
}

