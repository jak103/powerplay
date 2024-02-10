package auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/jak103/powerplay/internal/config"
	"github.com/jak103/powerplay/internal/utils/log"
)

func GenerateJwt(keyId int) (string, error) {
	mySigningKey := []byte(config.Vars.JwtSecret)
	now := time.Now()
	expiration := now.AddDate(0, 0, 1)
	// Create the Claims
	claims := &jwt.RegisteredClaims{
		ID:        fmt.Sprintf("%v", keyId),
		IssuedAt:  jwt.NewNumericDate(now),
		ExpiresAt: jwt.NewNumericDate(expiration),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwt, err := token.SignedString(mySigningKey)
	if err != nil {
		log.WithErr(err).Alert("Failed to sign JWT")
		return "", err
	}

	return jwt, nil
}
