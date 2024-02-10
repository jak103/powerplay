package schedule

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jak103/powerplay/internal/server/apis"
	"github.com/jak103/powerplay/internal/server/services/auth"
	"github.com/jak103/powerplay/internal/utils/responder"
)

func init() {
	apis.RegisterHandler(fiber.MethodPost, "/rsvp", auth.Authenticated, handleRsvp)
}

func handleRsvp(c *fiber.Ctx) error {

	return responder.NotYetImplemented(c)
}
