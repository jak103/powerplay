package schedule

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/jak103/powerplay/internal/db"
	"github.com/jak103/powerplay/internal/server/apis"
	"github.com/jak103/powerplay/internal/server/apis/schedule/pkg/analysis"
	"github.com/jak103/powerplay/internal/server/apis/schedule/pkg/models"
	"github.com/jak103/powerplay/internal/server/apis/schedule/pkg/parser"
	"github.com/jak103/powerplay/internal/server/services/auth"
	"github.com/jak103/powerplay/internal/utils/log"
	"github.com/jak103/powerplay/internal/utils/responder"
)

type response struct {
	TeamStats []models.TeamStats
}

func init() {
	apis.RegisterHandler(fiber.MethodGet, "/schedule/analysis/:seasonId<int/>", auth.Public, handleAnalysis)
}

func handleAnalysis(c *fiber.Ctx) error {
	//TODO: get the schedule from the database with this id
    seasonId, err := strconv.Atoi(c.Params("seasonId"))
    if err != nil {
        log.Error("Season Id was not an integer", err)
        responder.BadRequest(c, "Bad Request: Season Id was not an integer")
    }

    db := db.GetSession(c)
    games, err := db.GetGamesBySeason(seasonId)
    if err != nil {
        log.Error("Could not retrieve the games from the database")
        responder.InternalServerError(c, "Could not retrieve the games from the database")
    }

    //TODO: Change this to use different model.Game
	_, ts := analysis.RunTimeAnalysis(games)

	data := response{
		TeamStats: analysis.Serialize(ts),
	}

	return responder.OkWithData(c, data)
}

func printTeamSchedules(games []models.Game, seasonConfig models.SeasonConfig) {
	for _, league := range seasonConfig.Leagues {
		for _, team := range league.Teams {
			log.Info("-----------\n%v\n", team.Name)
			for _, game := range games {
				if team.Name == game.Team1Name || team.Name == game.Team2Name {
					log.Info("%s\n", game)
				}
			}
		}
	}
}
