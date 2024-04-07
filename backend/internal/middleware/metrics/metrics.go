package metrics

import "github.com/gofiber/fiber/v2"

func New() fiber.Handler {
	return middleware
}

func middleware(c *fiber.Ctx) error {
	// TODO write metrics middleware
	// TODO have the metrics include the Server-Timing header https://developer.mozilla.org/en-US/docs/Web/API/Performance_API/Server_timing
	return c.Next()
}
