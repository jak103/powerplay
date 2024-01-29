package auth

import "github.com/gofiber/fiber/v2"

type Type string

const (
	NONE Type = "none"
	JWT  Type = "jwt"
)

func GetMiddleware() fiber.Handler {
	// Need to decided between this implementations
	// TODO key auth https://github.com/gofiber/fiber/tree/v2/middleware/keyauth
	// TODO https://github.com/gofiber/contrib/tree/main/jwt
	return nil
}
