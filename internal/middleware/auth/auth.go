package auth

import (
	"errors"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jak103/powerplay/internal/config"
	"github.com/jak103/powerplay/internal/db"
	"github.com/jak103/powerplay/internal/server/apis"
	"github.com/jak103/powerplay/internal/server/services/auth"
	"github.com/jak103/powerplay/internal/utils/locals"
	"github.com/jak103/powerplay/internal/utils/log"
	"github.com/jak103/powerplay/internal/utils/responder"
)

func NewKeyRecord() fiber.Handler {
	return fetchKeyRecord
}

func NewAuthorizer() fiber.Handler {
	return authorizeRequest
}

func fetchKeyRecord(c *fiber.Ctx) error {
	// Get the JWT
	key := c.Cookies("Authorization")
	if len(key) == 0 {
		return c.Next() // No key record here, moving on
	}

	claims := jwt.RegisteredClaims{}
	key = strings.TrimPrefix(key, "Bearer ")

	token, err := jwt.ParseWithClaims(key, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Vars.JwtSecret), nil
	})

	// If the JWT is present, fetch the keyrecord
	if token.Valid {
		db := db.GetSession(c)
		id, _ := strconv.Atoi(claims.ID)
		record, err := db.GetKeyRecordById(uint(id))
		if err != nil {
			log.WithErr(err).Alert("Someone has had a valid JWT, but failed to get the keyrecord (ID: %v) from the DB", claims.ID)
			return responder.InternalServerError(c)
		}

		if record == nil {
			return responder.BadRequest(c, "JWT is no longer valid")
		}

		// store keyrecord in locals
		locals.SetKeyRecord(c, *record)

		return c.Next() // Got the keyrecord, proceeding with call
	}

	switch {
	case errors.Is(err, jwt.ErrTokenMalformed):
		return responder.BadRequest(c, jwt.ErrTokenMalformed.Error())
	case errors.Is(err, jwt.ErrTokenSignatureInvalid):
		// Invalid signature
		return responder.BadRequest(c, "Invalid JWT")
	case errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet):
		return responder.BadRequest(c, "Expired JWT")
	default:
		log.WithErr(err).Alert("Unhandled error while fetching keyrecord")
		return responder.InternalServerError(c)
	}
}

func authorizeRequest(c *fiber.Ctx) error {
	record := locals.KeyRecord(c)
	// two tiers of checks

	// 1) Is there a key, aka are you logged in
	if record == nil {
		return responder.Unauthorized(c, "Not logged in")
	}

	// 2) Do you have the right role?
	authorizedRoles := apis.GetRole(c.Route().Method, c.Route().Path)
	if auth.HasCorrectRole(record.Roles, authorizedRoles) {
		return c.Next() // User has correct role, let them through
	}

	return responder.Forbidden(c) // They do not have the right role
}
