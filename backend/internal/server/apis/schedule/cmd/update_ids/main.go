package main

import (
	"hockey/pkg/csv"
	"hockey/pkg/parser"
)

func main() {
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
}
