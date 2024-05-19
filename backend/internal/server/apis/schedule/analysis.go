package schedule

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jak103/powerplay/internal/server/apis"
	"github.com/jak103/powerplay/internal/server/apis/schedule/pkg/analysis"
	"github.com/jak103/powerplay/internal/server/apis/schedule/pkg/models"
	"github.com/jak103/powerplay/internal/server/apis/schedule/pkg/parser"
	"github.com/jak103/powerplay/internal/server/services/auth"
	"github.com/jak103/powerplay/internal/utils/log"
	"github.com/jak103/powerplay/internal/utils/responder"
)

type response struct {
	TeamStats []models.TeamData `json:"teamStats"`
}

func init() {
	apis.RegisterHandler(fiber.MethodGet, "/schedule/analysis/:scheduleId<int/>", auth.ManagerOnly, handleAnalysis)
}

func handleAnalysis(c *fiber.Ctx) error {
	games, seasonConfig := parser.ReadGames("summer_2024")

	scheduleId := c.Params("scheduleId")
	//TODO: get the schedule from the database with this id
	log.Info(scheduleId)

	// TODO: get team stats and serialize all of the teams stats
	// Research if we are storing these in a database, if not, we can store them
	// The map contains the team name, an example json object might look like this:
	// {
	//     teamName1: {
	//         team1Data
	//     }
	//     teamName2: {
	//         team2Data
	//     }
	// }
	_, ts := analysis.RunTimeAnalysis(games)

	printTeamSchedules(games, seasonConfig)

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
