package goals

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jak103/powerplay/internal/server/apis"
	"github.com/jak103/powerplay/internal/server/services/auth"
	"github.com/jak103/powerplay/internal/utils/responder"

)


func init() {
	apis.RegisterHandler(fiber.MethodPost, "/goals/submit", auth.Public, postGoalsHandler)
}

func postGoalsHandler(c *fiber.Ctx) error {
	// If valid structure in post request, continue on

	// Connect to database and insert goal

	return responder.NotYetImplemented(c)

}
