package stats

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/jak103/powerplay/internal/db"
	"github.com/jak103/powerplay/internal/models"
	"github.com/jak103/powerplay/internal/server/apis"
	"github.com/jak103/powerplay/internal/server/services/auth"
	"github.com/jak103/powerplay/internal/utils/locals"
	"github.com/jak103/powerplay/internal/utils/responder"
)

func init() {
	//apis.RegisterHandler(fiber.MethodGet, "/penaltyTypes", auth.Public, getPenaltyTypes)
	apis.RegisterHandler(fiber.MethodGet, "/penalties", auth.Public, getPenaltiesHandler)
	apis.RegisterHandler(fiber.MethodPost, "/penalties", auth.Public, postPenaltyHandler)
}

func getPenaltiesHandler(c *fiber.Ctx) error {
	log := locals.Logger(c)
	log.Info("Handling getting all penalties")
	db := db.GetSession(c)
	penalties, err := db.GetPenalties()
	if err != nil {
		log.WithErr(err).Alert("Failed to get all penalties from the database")
		return err
	}

	jsonData, err := json.Marshal(penalties)
	if err != nil {
		log.WithErr(err).Alert("Failed to serialize penalties response payload")
		return err
	}

	c.Type("json")

	// Send JSON response
	return c.Send(jsonData)
}

func postPenaltyHandler(c *fiber.Ctx) error {
	log := locals.Logger(c)
	log.Info("Handling creating new penalty")
	log.Debug("body: %q", c.Request().Body())

	// Parse penalty
	penaltyRequest := &models.Penalty{}
	err := c.BodyParser(penaltyRequest)
	if err != nil {
		log.WithErr(err).Alert("Failed to parse penalty request payload")
		return responder.InternalServerError(c)
	}

	db := db.GetSession(c)
	err = db.CreatePenalty(penaltyRequest)
	if err != nil {
		log.WithErr(err).Alert("Failed to save penalty request")
		return responder.InternalServerError(c)
	}

	return responder.Ok(c)
}
