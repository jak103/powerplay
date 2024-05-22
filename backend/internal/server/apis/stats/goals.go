package stats

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jak103/powerplay/internal/server/apis"
	"github.com/jak103/powerplay/internal/server/services/auth"
	"github.com/jak103/powerplay/internal/utils/responder"
	"github.com/jak103/powerplay/internal/models"
	"github.com/jak103/powerplay/internal/db"
	"github.com/jak103/powerplay/internal/utils/locals"
)


func init() {
	apis.RegisterHandler(fiber.MethodPost, "/goals", auth.Public, postGoalsHandler)
}

func postGoalsHandler(c *fiber.Ctx) error {
	log := locals.Logger(c)
	log.Info("Handling creating new goal")
	log.Debug("body: %q", c.Request().Body())
	goalPostRequest := new(models.Goal)
	err := c.BodyParser(goalPostRequest)
	
	// If valid structure in post request, continue on
	if err != nil{
		log.WithErr(err).Error("Failed to parse Goal POST request.")
		return err
	}

	// Connect to database and insert goal
	db := db.GetSession(c)
	record, err := db.SaveGoal(goalPostRequest)
	
	if err != nil{
		log.WithErr(err).Alert("Failed to parse goal request payload")
		return responder.InternalServerError(c)
	}

	if record == nil {
		return responder.BadRequest(c, "Could not post goal into database")
	}
	
	return responder.Ok(c)

}