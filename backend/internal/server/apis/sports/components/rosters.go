package components

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jak103/powerplay/internal/db"
	"github.com/jak103/powerplay/internal/models"
	"github.com/jak103/powerplay/internal/server/apis"
	"github.com/jak103/powerplay/internal/server/services/auth"
	"github.com/jak103/powerplay/internal/utils/locals"
	"github.com/jak103/powerplay/internal/utils/log"
	"github.com/jak103/powerplay/internal/utils/responder"
)

func init() {
	apis.RegisterHandler(fiber.MethodPost, "/rosters", auth.Authenticated, postRoster)
	apis.RegisterHandler(fiber.MethodGet, "/rosters", auth.Public, getRosters)
}

func postRoster(c *fiber.Ctx) error {
	// need to setup middleware for auth...

	// Get user that is hitting api
	// Validate that they are authenticated to hit this endpoint

	type RosterRequest struct {
		CaptainID uint   `json:"captain_id"`
		PlayerIDs []uint `json:"player_ids"`
	}

	log := locals.Logger(c)

	body := &RosterRequest{}
	err := c.BodyParser(body)
	if err != nil {
		msg := "Failed to parse roster request payload"
		log.WithErr(err).Alert(msg)
		return responder.BadRequest(c, msg)
	}

	db := db.GetSession(c)

	log.Debug("Captain : %d", body.CaptainID)
	capt, err := db.GetUserByID(body.CaptainID)
	if err != nil {
		log.WithErr(err).Alert("Failed to get captain")
		return responder.InternalServerError(c)
	}

	if capt == nil {
		log.WithErr(err).Alert("Captain DNE")
		return responder.InternalServerError(c)
	}

	log.Debug("Players : %v", body.PlayerIDs)
	players, err := db.GetUsersByIDs(body.PlayerIDs)
	if err != nil {
		log.WithErr(err).Alert("Failed to get players")
		return responder.InternalServerError(c)
	}

	roster := models.Roster{
		CaptainID: capt.ID,
		Captain:   *capt,
		Players:   players,
	}

	_, err = db.CreateRoster(&roster)
	if err != nil {
		log.WithErr(err).Alert("Failed to save roster request")
		return responder.InternalServerError(c)

	}

	return responder.Ok(c)
}

func getRosters(c *fiber.Ctx) error {

	db := db.GetSession(c)
	rosters, err := db.GetRosters()
	if err != nil {
		log.WithErr(err).Alert("Failed to get all rosters from the database")
		return err
	}

	return responder.OkWithData(c, rosters)
}
