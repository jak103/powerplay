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
	apis.RegisterHandler(fiber.MethodPost, "/shotsongoal", auth.Public, postShotsOnGoalHandler)
	// Todo add register handler for GET
}

func postShotsOnGoalHandler(c *fiber.Ctx) error {
	log := locals.Logger(c)
	log.Info("Handling creating new shot on goal")
	log.Debug("body: %q", c.Request().Body())
	shotOnGoalRequest := new(models.ShotOnGoal)
	err := c.BodyParser(shotOnGoalRequest)
	if err != nil {
		log.WithErr(err).Error("Failed to parse ShotsOnGoal POST request.")
		return err
	}

	db := db.GetSession(c)
	record, err := db.SaveShotOnGoal(shotOnGoalRequest)

	if err != nil {
		log.WithErr(err).Alert("Failed to parse request payload")
		return responder.InternalServerError(c)
	}

	if record == nil {
		return responder.BadRequest(c,"Could not Post shot on goal to database.")
	}
	return responder.Ok(c)	
}