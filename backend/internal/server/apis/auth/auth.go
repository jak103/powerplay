package auth

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jak103/powerplay/internal/config"
	"github.com/jak103/powerplay/internal/server/apis"
	"github.com/jak103/powerplay/internal/server/services/auth"
	"github.com/jak103/powerplay/internal/utils/locals"
	"github.com/jak103/powerplay/internal/utils/log"
	"github.com/jak103/powerplay/internal/utils/responder"
)

type request struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type response struct {
	Jwt        string    `json:"jwt"`
	Expiration time.Time `json:"expiration"`
}

func init() {
	apis.RegisterHandler(fiber.MethodPost, "/auth", auth.Public, postAuthHandler)
}

func postAuthHandler(c *fiber.Ctx) error {
	log := locals.Logger(c)
	creds := request{}
	err := c.BodyParser(&creds)
	if err != nil {
		log.WithErr(err).Error("Failed to parse authentication credentials")
		return responder.BadRequest(c, "Failed to parse authentication credentials")
	}

	// TODO look up user in database

	jwt, err := generateJwt(1)
	if err != nil {
		log.WithErr(err).Alert("Failed to generate JWT")
	}

	token := response{
		Jwt:        jwt,
		Expiration: time.Now().Add(24 * time.Hour),
	}

	return responder.OkWithData(c, token)
}

func generateJwt(keyId int) (string, error) {
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
