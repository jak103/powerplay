package metrics

import "github.com/gofiber/fiber/v2"

func New() fiber.Handler {
	return middleware
}

func middleware(c *fiber.Ctx) error {
	// TODO write metrics middleware
	return c.Next()
}
