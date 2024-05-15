package main

import (
	"fmt"
	"hockey/pkg/csv"
	"hockey/pkg/parser"
)

type RefScheduleRow struct {
	Start           string `csv:"Start Date and Time"`
	DurationHours   string `csv:"Duration Hours"`
	DurationMinutes string `csv:"Duration Minutes"`
	Location        string `csv:"Arena/Rink"`
	Level           string `csv:"Game Level"`
	Home            string `csv:"Home Team"`
	Away            string `csv:"Away Team"`
}

func main() {
	games, seasonConfig := parser.ReadGames("spring_2024")

	refSchedule := make([]RefScheduleRow, 0)

	for i, game := range games {
		for _, league := range seasonConfig.Leagues {
			for _, team := range league.Teams {
				if game.Team1Name == team.Name {
					games[i].League = league.Name
					break
				}
			}
		}

		row := RefScheduleRow{
			Start:           game.Start.Format("1/2/06 3:04 PM"),
			DurationHours:   "1",
			DurationMinutes: "15",
			Location:        "George S. Eccles Ice Center",
			Level:           fmt.Sprintf("%s League", game.League),
			Home:            game.Team1Name,
			Away:            game.Team2Name,
		}

		refSchedule = append(refSchedule, row)
	}

	csv.GenerateCsv(refSchedule, "ref_schedule.csv")
}
