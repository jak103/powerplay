package analysis

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jak103/powerplay/internal/models"
	"github.com/jak103/powerplay/internal/server/apis"
	"github.com/jak103/powerplay/internal/server/apis/schedule/helpers/analysis"
	"github.com/jak103/powerplay/internal/server/apis/schedule/helpers/read"
	"github.com/jak103/powerplay/internal/server/services/auth"
	"github.com/jak103/powerplay/internal/utils/log"
	"github.com/jak103/powerplay/internal/utils/responder"
)

func init() {
	apis.RegisterHandler(fiber.MethodPost, "/schedule/analysis", auth.Authenticated, handleAnalysis)
}

func handleAnalysis(c *fiber.Ctx) error {
	games, seasonConfig := read.Games("test")

	// TODO numberOfGamesPerTeam is hardcoded to 10 for now. We need to read this from the ice_times.csv file.
	analysis.RunTimeAnalysis(games, 10)

	printTeamSchedules(games, seasonConfig)
	return responder.NotYetImplemented(c)
}

func printTeamSchedules(games []models.Game, seasonConfig models.SeasonConfig) {
	for _, league := range seasonConfig.Leagues {
		for _, team := range league.Teams {
			log.Info("-----------\n%v\n", team.Name)
			for _, game := range games {
				if team.Name == game.Teams[0].Name || team.Name == game.Teams[1].Name {
					log.Info("%s\n", game)
				}
			}
		}
	}
}
