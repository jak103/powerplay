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

func NewRequestingUser() fiber.Handler {
	return fetchUser
}

func NewAuthorizer() fiber.Handler {
	return authorizeRequest
}

func fetchUser(c *fiber.Ctx) error {
	// Get the JWT
	key := c.Cookies("Authorization")
	if len(key) == 0 {
		return c.Next() // No user here, moving on
	}

	claims := jwt.RegisteredClaims{}
	key = strings.TrimPrefix(key, "Bearer ")

	token, err := jwt.ParseWithClaims(key, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Vars.JwtSecret), nil
	})

	// If the JWT is present, fetch the user
	if token.Valid {
		db := db.GetSession(c)
		id, _ := strconv.Atoi(claims.ID)
		user, err := db.GetUserById(id)
		if err != nil {
			log.WithErr(err).Alert("Someone has had a valid JWT, but failed to get the user (ID: %v) from the DB", claims.ID)
			return responder.InternalServerError(c)
		}

		if user == nil {
			return responder.BadRequest(c, "JWT is no longer valid")
		}

		// store user in locals
		locals.SetRequestingUser(c, *user)

		return c.Next() // Got the user, proceeding with call
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
		log.WithErr(err).Alert("Unhandled error while fetching user")
		return responder.InternalServerError(c)
	}
}

func authorizeRequest(c *fiber.Ctx) error {
	user := locals.RequestingUser(c)
	// two tiers of checks

	// 1) Is there a key, aka are you logged in
	if user == nil {
		return responder.Unauthorized(c, "Not logged in")
	}

	// 2) Do you have the right role?
	authorizedRoles := apis.GetRole(c.Route().Method, c.Route().Path)
	if auth.HasCorrectRole(user.Roles, authorizedRoles) {
		return c.Next() // User has correct role, let them through
	}

	return responder.Forbidden(c) // They do not have the right role
}
