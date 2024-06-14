package stats

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
	apis.RegisterHandler(fiber.MethodPost, "/goals", auth.Public, postGoalsHandler)
	apis.RegisterHandler(fiber.MethodGet, "/goals", auth.Public, getGoalsHandler)
}

func postGoalsHandler(c *fiber.Ctx) error {
	log := locals.Logger(c)
	log.Debug("body: %q", c.Request().Body())
	goalPostRequest := new(models.Goal)
	err := c.BodyParser(goalPostRequest)

	// If valid structure in post request, continue on
	if err != nil {
		log.WithErr(err).Error("Failed to parse Goal POST request.")
		return responder.BadRequest(c)
	}

	// Connect to database and insert goal
	db := database.newSession(c)
	record, err := db.SaveGoal(goalPostRequest)

	if err != nil {
		log.WithErr(err).Alert("Failed to parse goal request payload")
		return responder.InternalServerError(c)
	}

	if record == nil {
		return responder.BadRequest(c, "Could not post goal into database")
	}

	return responder.Ok(c)

}

func getGoalsHandler(c *fiber.Ctx) error {
	log := locals.Logger(c)
	db := db.GetSession(c)
	goals, err := db.GetGoals()
	if err != nil {
		log.WithErr(err).Alert("Failed to get all goals from the database")
		return err
	}
	// Send JSON response
	return responder.OkWithData(c, goals)
}
