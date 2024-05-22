package rsvp

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/jak103/powerplay/internal/server/apis"
	"github.com/jak103/powerplay/internal/server/services/auth"
	"github.com/jak103/powerplay/internal/utils/responder"
)

func init() {
	apis.RegisterHandler(fiber.MethodPost, "/rsvp", auth.Authenticated, handleRsvp)
}

type BodyDto struct {
	SeasonId int `json:"season_id"`
	LeagueId int `json:"league_id"`
	TeamId   int `json:"team_id"`
	GameId   int `json:"game_id"`
	Rsvp     int `json:"rsvp"`
}

func handleRsvp(c *fiber.Ctx) error {
	// TODO I think the body of the request should be a JSON object with the following fields:
	// - season_id (int)
	// - league_id (int)
	// - team_id (int)
	// - game_id (int)
	// - rsvp (string) - "yes":1 or "no":0
	// With this information, we can update the team roster for that game.
	// We should also check if the user is on the team roster for that game.
	// If not, we should check if they are a sub for that team.
	// So we need to have another table in the database that keeps track of who played in each game.
	return responder.NotYetImplemented(c)
}

func readBody(c *fiber.Ctx) (BodyDto, error) {
	body := c.Body()
	var bodyDto BodyDto
	err := json.Unmarshal(body, &bodyDto)
	if err != nil {
		return BodyDto{}, responder.BadRequest(c, "Error reading body")
	}
	return bodyDto, nil
}
