package auth

import "github.com/gofiber/fiber/v2"

type Type string

const (
	NONE Type = "none"
	JWT  Type = "jwt"
)

func GetMiddleware() fiber.Handler {
	return nil
}
