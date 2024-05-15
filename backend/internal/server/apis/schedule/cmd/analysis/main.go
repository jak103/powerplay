package main

import (
	"fmt"
	"hockey/pkg/analysis"
	"hockey/pkg/models"
	"hockey/pkg/parser"
)

func main() {
	games, seasonConfig := parser.ReadGames("summer_2024")

	analysis.RunTimeAnalysis(games)

	printTeamSchedules(games, seasonConfig)
}

func printTeamSchedules(games []models.Game, seasonConfig models.SeasonConfig) {
	for _, league := range seasonConfig.Leagues {
		for _, team := range league.Teams {
			fmt.Printf("-----------\n%v\n", team.Name)
			for _, game := range games {
				if team.Name == game.Team1Name || team.Name == game.Team2Name {
					fmt.Printf("%s\n", game)
				}
			}
		}
	}
}
