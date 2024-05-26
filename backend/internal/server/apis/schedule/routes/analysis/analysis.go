package analysis

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/jak103/powerplay/internal/models"
	"github.com/jak103/powerplay/internal/server/apis"
	"github.com/jak103/powerplay/internal/server/apis/schedule/helpers/analysis"
	"github.com/jak103/powerplay/internal/server/apis/schedule/helpers/read"
	"github.com/jak103/powerplay/internal/server/apis/schedule/helpers/util"
	"github.com/jak103/powerplay/internal/server/services/auth"
	"github.com/jak103/powerplay/internal/utils/log"
	"github.com/jak103/powerplay/internal/utils/responder"
	"time"
)

func init() {
	apis.RegisterHandler(fiber.MethodPost, "/schedule/analysis", auth.Authenticated, handleAnalysis)
}

func handleAnalysis(c *fiber.Ctx) error {
	games, leagues, err := something(c, "test")
	if err != nil {
		log.Error("Error reading games and leagues: %v\n", err)
	}

	// TODO numberOfGamesPerTeam is hardcoded to 10 for now. We need to read this from the upload.csv file.
	analysis.RunTimeAnalysis(games, 10)

	printTeamSchedules(games, leagues)
	return responder.NotYetImplemented(c)
}

func something(c *fiber.Ctx, season string) ([]models.Game, []models.League, error) {
	// TODO I dont know the purpose of this function
	// it seems like this function isn't needed anymore since we can just call read.Games and read.Leagues directly.
	// This function shouldn't be modifying the games and leagues as we are wanting to analyze the data as is.

	if c == nil || len(season) == 0 {
		return nil, nil, errors.New("invalid uploads")
	}

	games, err := read.Games(c, season)
	if err != nil {
		return nil, nil, err
	}

	leagues, err := read.Leagues(c, season)
	if err != nil {
		return nil, nil, err
	}

	for i := range games {
		games[i].Start, err = time.Parse("01/02/2006 15:04", fmt.Sprintf("%v %v", games[i].StartDate, games[i].StartTime))
		if err != nil {
			log.Error("Time parse error: %v\n", err)
			return nil, nil, errors.New("time parse error")
		}

		games[i].IsEarly = util.IsEarlyGame(games[i].Start.Hour(), games[i].Start.Minute())
		for _, league := range leagues {
			for _, team := range league.Teams {
				if games[i].Teams[0].Name == team.Name {
					games[i].League = league.Name
					break
				}
			}
		}
	}

	return games, leagues, nil
}

func printTeamSchedules(games []models.Game, leagues []models.League) {
	for _, league := range leagues {
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
