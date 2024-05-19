package schedule

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jak103/powerplay/internal/server/apis/schedule/pkg/csv"
	"github.com/jak103/powerplay/internal/server/apis/schedule/pkg/parser"
	"github.com/jak103/powerplay/internal/utils/responder"
)

func handleUpdateIds(c *fiber.Ctx) error {
	games, seasonConfig := parser.ReadGames("spring_2024")

	for gameIndex, game := range games {
		for _, league := range seasonConfig.Leagues {
			if game.League == league.Name {
				for _, team := range league.Teams {
					if team.Name == game.Team1Name {
						games[gameIndex].Team1Id = team.Id
					}

					if team.Name == game.Team2Name {
						games[gameIndex].Team2Id = team.Id
					}
				}
			}
		}
	}

	csv.GenerateCsv(games, "schedule_updated.csv")
	return responder.NotYetImplemented(c)
}
