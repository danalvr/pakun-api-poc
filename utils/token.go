package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("secret")

func GenerateJWT(identifier string) (string, error) {
	claims := jwt.MapClaims{
		"identifier": identifier,
		"exp": time.Now().Add(time.Hour * 1).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}