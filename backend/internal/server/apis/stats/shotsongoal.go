package stats

import (
	"github.com/jak103/powerplay/internal/server/apis"
	"github.com/jak103/powerplay/internal/utils/locals"
	"github.com/jak103/powerplay/internal/utils/responder"
	"github.com/gofiber/fiber/v2"
	"github.com/jak103/powerplay/internal/db"
	"github.com/jak103/powerplay/internal/server/services/auth"
	"github.com/jak103/powerplay/internal/models"
)

func init() {
	apis.RegisterHandler(fiber.MethodPost, "shotsOnGoal", auth.Public, postShotsOnGoalHandler)
	// Todo add register handler for GET
}

func postShotsOnGoalHandler(c *fiber.Ctx) error {
	log := locals.Logger(c)
	shotOnGoal := new(models.ShotOnGoal)
	err := c.BodyParser(shotOnGoal)
	if err != nil {
		log.WithErr(err).Alert("Failed to parse ShotsOnGoal POST request.")
		return err
	}

	db := db.GetSession(c)
	model, err := db.SaveShotOnGoal(shotOnGoal)

	if err != nil {
		log.WithErr(err).Alert("Failed to ")
		return responder.InternalServerError(c)
	}

	if model == nil {
		return responder.InternalServerError(c)
	}
	return responder.Ok(c)	
}