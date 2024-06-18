package player

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/jak103/powerplay/internal/db"
	"github.com/jak103/powerplay/internal/server/apis"
	"github.com/jak103/powerplay/internal/server/services/auth"
	"github.com/jak103/powerplay/internal/utils/responder"
)

func init() {
	apis.RegisterHandler(fiber.MethodPost, "/rsvp/player", auth.Authenticated, handleRsvp)
}

type Body struct {
    UserId   uint `json:"user_id"`
	TeamId   uint `json:"team_id"`
	GameId   uint `json:"game_id"`
	Rsvp     int  `json:"player"`
}

const (
    NotAttending = iota
    Attending
)

func handleRsvp(c *fiber.Ctx) error {
	// TODO I think the body of the request should be a JSON object with the following fields:
	// - team_id (int)
	// - game_id (int) - Should only need game_id and team_id, the rest can be derived
	// - player (string) - "yes":1 or "no":0
	// With this information, we can update the team roster for that games.
	// We should also check if the user is on the team roster for that games.
	// If not, we should check if they are a sub for that team.
	// So we need to have another table in the database that keeps track of who played in each games.
    body, err := readBody(c)
    if err != nil {
        return responder.InternalServerError(c, "Could not read the body of the request")
    }

    //If they are attending, then put them on the roster
    if body.Rsvp == Attending {
        session := db.GetSession(c)
        team, err := session.GetTeamByID(string(body.TeamId))
        if err != nil {
            return responder.InternalServerError(c, err.Error())
        }
        game, err := session.GetGameById(body.GameId)
        if err != nil {
            return responder.InternalServerError(c, err.Error())
        }

        user, err := session.GetUserById(body.UserId)
        if err != nil {
            return responder.InternalServerError(c, err.Error())
        }

        // Check if the team they are rsvp'ing for is home or away team
        if team == &game.HomeTeam {
            _ = append(game.HomeTeamRoster.Players, user)
            session.SaveGame(*game)
        }

        if team == &game.AwayTeam {
            _ = append(game.HomeTeamRoster.Players, user)
            session.SaveGame(*game)
        }
    }

	return responder.Ok(c, "Successfully rsvp'd")
}

func readBody(c *fiber.Ctx) (Body, error) {
	dto := struct {
        UserId             uint   `json:"user_id"`
		TeamId             uint   `json:"team_id"`
		GameId             uint   `json:"game_id"`
		Rsvp               int    `json:"rsvp"`
	}{}

	if err := c.BodyParser(&dto); err != nil {
		return Body{}, responder.BadRequest(c, "Failed to parse body of request")
	}

	body := Body{
        UserId: dto.UserId,
		TeamId: dto.TeamId,
        GameId: dto.GameId,
		Rsvp:   dto.Rsvp,
	}
	return body, nil
}
