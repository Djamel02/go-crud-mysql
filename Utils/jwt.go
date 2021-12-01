package utils

import (
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateToken(id int) (*string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(id),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})

	// Create Our token
	token, err := claims.SignedString([]byte(GetEnvironmentVars("JWT_SECRET")))
	if err != nil {
		return nil, err
	}

	return &token, nil
}
