package parser

import (
	"fmt"
	"github.com/jak103/powerplay/internal/server/apis/schedule/pkg/models"
	"os"
	"time"

	"github.com/gocarina/gocsv"
)

func ReadGames(season string) ([]models.Game, models.SeasonConfig) {
	fmt.Println("Reading config file season_config.yml")

	seasonConfig, err := SeasonConfig(season)
	if err != nil {
		fmt.Println("Error reading file", err)
	}

	games := []models.Game{}

	scheduleFile, err := os.OpenFile(fmt.Sprintf("schedule_%s.csv", season), os.O_RDONLY, os.ModePerm)
	if err != nil {
		fmt.Println("Failed to open schedule file", err)
	}

	err = gocsv.UnmarshalFile(scheduleFile, &games)
	if err != nil {
		fmt.Println("Failed to unmarshal schedule file", err)
	}

	fmt.Printf("Read %v games\n", len(games))

	for i := range games {
		games[i].Start, err = time.Parse("01/02/2006 15:04", fmt.Sprintf("%v %v", games[i].StartDate, games[i].StartTime))
		if err != nil {
			fmt.Println("Time parse error", err)
		}

		switch games[i].Start.Hour() {
		case 20:
			games[i].IsEarly = true
		case 21:
			if games[i].Start.Minute() <= 15 {
				games[i].IsEarly = true
			} else {
				games[i].IsEarly = false
			}

		case 22, 23:
			games[i].IsEarly = false
		}

		for _, league := range seasonConfig.Leagues {
			for _, team := range league.Teams {
				if games[i].Team1Name == team.Name {
					games[i].League = league.Name
					break
				}
			}
		}
	}

	return games, *seasonConfig
}
