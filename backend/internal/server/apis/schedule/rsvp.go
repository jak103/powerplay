package schedule

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jak103/powerplay/internal/utils/responder"
)

func handleRsvp(c *fiber.Ctx) error {
	// TODO I think the body of the request should be a JSON object with the following fields:
	// - season_id (int)
	// - league_id (int)
	// - team_id (int)
	// - ice_time_id (int)
	// - rsvp (string) - "yes" or "no"
	// With this information, we can update the team roster for that game.
	// We should also check if the user is on the team roster for that game.
	//If not, we should check if they are a sub for that team.
	// So we need to have another table in the database that keeps track of who played in each game.
	return responder.NotYetImplemented(c)
}
