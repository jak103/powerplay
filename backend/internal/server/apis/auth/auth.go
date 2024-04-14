package auth

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jak103/powerplay/internal/config"
	"github.com/jak103/powerplay/internal/models"
	"github.com/jak103/powerplay/internal/server/apis"
	"github.com/jak103/powerplay/internal/server/services/auth"
	"github.com/jak103/powerplay/internal/utils/locals"
	"github.com/jak103/powerplay/internal/utils/log"
	"github.com/jak103/powerplay/internal/utils/responder"
)

type response struct {
	Jwt        string    `json:"jwt"`
	Expiration time.Time `json:"expiration"`
}

func init() {
	apis.RegisterHandler(fiber.MethodPost, "/auth", auth.Public, postAuthHandler)
}

func postAuthHandler(c *fiber.Ctx) error {
	log := locals.Logger(c)
	db := newSession(c)

	creds := models.Credentials{}
	err := c.BodyParser(&creds)
	if err != nil {
		log.WithErr(err).Error("Failed to parse authentication credentials")
		return responder.BadRequest(c, "Failed to parse authentication credentials")
	}
	log.Debug("Creds: %v", creds)
	user, err := db.GetUserByEmail(creds.Email)
	if err != nil {
		log.WithErr(err).Error("Failed to get user by email")
		return responder.InternalServerError(c)
	}

	if user == nil {
		log.Info("User with email %s doesn't exist", creds.Email)
		return responder.Unauthorized(c, "Invalid credentials")
	}

	log.Debug("Got user %s", user.Email)
	jwt, err := generateJwt(user.ID)
	if err != nil {
		log.WithErr(err).Alert("Failed to generate JWT")
	}

	token := response{
		Jwt:        jwt,
		Expiration: time.Now().Add(24 * time.Hour),
	}

	return responder.OkWithData(c, token)
}

func generateJwt(keyId uint) (string, error) {
	mySigningKey := []byte(config.Vars.JwtSecret)
	now := time.Now()
	expiration := now.AddDate(0, 0, 7) // Token is good for 1 week

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
