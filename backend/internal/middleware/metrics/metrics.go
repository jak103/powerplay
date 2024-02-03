package metrics

import "github.com/gofiber/fiber/v2"

func New() fiber.Handler {
	// TODO write metrics middleware
	return middleware
}

func middleware(c *fiber.Ctx) error {
	return c.Next()
}
