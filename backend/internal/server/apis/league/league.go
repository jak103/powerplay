package league

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jak103/powerplay/internal/db"
	"github.com/jak103/powerplay/internal/server/apis"
	"github.com/jak103/powerplay/internal/server/services/auth"
	"github.com/jak103/powerplay/internal/utils/locals"
	"github.com/jak103/powerplay/internal/utils/responder"
)

func init() {
	apis.RegisterHandler(fiber.MethodGet, "/leagues", auth.Public, getLeaguesHandler)
}

func getLeaguesHandler(c *fiber.Ctx) error {
	log := locals.Logger(c)
	db := db.GetSession(c)
	leagues, err := db.GetLeagues()
	if err != nil {
		// todo: create ticket to standardize this error message and pass in model name
		log.WithErr(err).Alert("Failed to get all leagues from the database")
		return err
	}

	return responder.OkWithData(c, leagues)
}
