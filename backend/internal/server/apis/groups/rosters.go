package groups

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
	apis.RegisterHandler(fiber.MethodPost, "/rosters", auth.Public, postRoster)
	apis.RegisterHandler(fiber.MethodGet, "/rosters", auth.Public, getRosters)
}

func postRoster(c *fiber.Ctx) error {
	type RosterRequest struct {
		CaptainEmail string   `json:"captain_email"`
		PlayerEmails []string `json:"player_emails"`
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

	log.Debug("Captain : %s", body.CaptainEmail)
	capt, err := db.GetUserByEmail(body.CaptainEmail)
	if err != nil {
		log.WithErr(err).Alert("Failed to get captain")
		return responder.InternalServerError(c)
	}

	if capt == nil {
		log.WithErr(err).Alert("Captain DNE")
		return responder.InternalServerError(c)
	}

	log.Debug("Players : %v", body.PlayerEmails)
	players, err := db.GetUserByEmails(body.PlayerEmails)
	if err != nil {
		log.WithErr(err).Alert("Failed to get players")
		return responder.InternalServerError(c)
	}

	roster := models.Roster{
		CaptainID: capt.ID,
		Captain:   *capt,
		Players:   players,
	}

	err = db.CreateRoster(&roster)
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
