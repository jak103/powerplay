package analysis

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

func init() {
	apis.RegisterHandler(fiber.MethodPost, "/schedule/analysis", auth.Authenticated, handleAnalysis)
}

func handleAnalysis(c *fiber.Ctx) error {
	games, seasonConfig := parser.ReadGames("summer_2024")

	analysis.RunTimeAnalysis(games)

	printTeamSchedules(games, seasonConfig)
	return responder.NotYetImplemented(c)
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
