package components

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jak103/powerplay/internal/db"
	"github.com/jak103/powerplay/internal/models"
	"github.com/jak103/powerplay/internal/server/apis"
	"github.com/jak103/powerplay/internal/server/services/auth"
	"github.com/jak103/powerplay/internal/utils/locals"
	"github.com/jak103/powerplay/internal/utils/responder"
)

func init() {
	apis.RegisterHandler(fiber.MethodGet, "/leagues", auth.Public, getLeaguesHandler)
	apis.RegisterHandler(fiber.MethodPost, "/leagues", auth.Public, postLeagueHandler)
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

func postLeagueHandler(c *fiber.Ctx) error {
	log := locals.Logger(c)

	leagueRequest := &models.League{}
	err := c.BodyParser(leagueRequest)
	if err != nil {
		log.WithErr(err).Alert("Failed to parse leagues request payload")
		return responder.BadRequest(c, "Failed to parse leagues request payload")
	}

	db := db.GetSession(c)
	err = db.CreateLeague(leagueRequest)
	if err != nil {
		log.WithErr(err).Alert("Failed to save leagues request")
		return responder.InternalServerError(c)
	}

	return responder.Ok(c)
}
