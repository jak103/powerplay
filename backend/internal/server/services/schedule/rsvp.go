package schedule

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jak103/powerplay/internal/auth"
	"github.com/jak103/powerplay/internal/server"
)

func init() {
	server.RegisterHandler(fiber.MethodPost, "/rsvp", auth.JWT, handleRsvp)
}

func handleRsvp(c *fiber.Ctx) error {

	return nil
}
