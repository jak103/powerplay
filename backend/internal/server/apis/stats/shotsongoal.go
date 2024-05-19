package stats

import (
	"github.com/jak103/powerplay/internal/server/apis"
	"github.com/jak103/powerplay/internal/utils/locals"
	"github.com/jak103/powerplay/internal/utils/log"
	"github.com/jak103/powerplay/internal/utils/responder"
)

func init() {
	apis.RegisterHandler(fiber.MethodPost, "shotsOnGoal", auth.Public, postShotsOnGoalHandler)
	// Todo add register handler for GET
}

func postShotsOnGoalHandler(c *fiber.Ctx) error {
	log := locals.Logger(c)
	shotOnGoal := new(models.ShotOnGoal)
	err := c.BodyParser(ShotOnGoal)
	if err != nil {
		log.WithErr(err).Error("Failed to parse ShotsOnGoal POST request.")
		return err
	}

	db := c.GetSession(c)
	model, err := db.SaveShotOnGoal(shotOnGoal)

	if err != nil {
		log.WithErr(err).Error("")
		return responder.internalServerError(c, )
	}

	if record == nil {
		log.Logger
		return responder.internalServerError(c, "failed to POST ShotOnGoal.")
	}
	return responder.Ok(c)	
}