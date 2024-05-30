package save

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/jak103/powerplay/internal/db"
	dbModels "github.com/jak103/powerplay/internal/models"
	scheduleModels "github.com/jak103/powerplay/internal/server/apis/schedule/pkg/models"
)

func ToDb(c *fiber.Ctx, games []scheduleModels.Game) error {
	if c == nil || len(games) == 0 {
		return errors.New("invalid input")
	}
	session := db.GetSession(c)
	dbGames := make([]dbModels.Game, len(games))
	for i, game := range games {
		dbGames[i] = mapScheduleGameToDbGame(game)
	}

	for _, dbGame := range dbGames {
		if err := session.Connection.Save(&dbGame).Error; err != nil {
			return err
		}
	}

	return nil
}

func mapScheduleGameToDbGame(game scheduleModels.Game) dbModels.Game {
	return dbModels.Game{
		HomeTeam: []dbModels.Team{
			Name: game.Team1Name
		},
		Start: game.Start,
		End:   game.End,
	}
}
