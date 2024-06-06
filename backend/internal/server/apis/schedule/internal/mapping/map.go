package mapping

import (
	"github.com/jak103/powerplay/internal/models"
	"github.com/jak103/powerplay/internal/server/apis/schedule/internal/structures"
)

func MapGameStructToGameModel(games []structures.Game) []models.Game {
	var gameModels []models.Game
	for _, game := range games {
		gameModels = append(gameModels, models.Game{
			DbModel:  game.DbModel,
			SeasonID: game.SeasonID,
			Start:    game.Start,
			Venue: models.Venue{
				Name:        "George S. Eccles Ice Center",
				Address:     "2825 N 200 E North Logan, UT 84341 United States",
				LockerRooms: []string{"1", "2", "3", "4", "5"},
			},
			VenueID:             0,
			Status:              models.SCHEDULED,
			HomeTeam:            game.HomeTeam,
			HomeTeamID:          game.HomeTeamID,
			HomeTeamRoster:      game.HomeTeamRoster,
			HomeTeamRosterID:    game.HomeTeamRosterID,
			HomeTeamLockerRoom:  game.HomeTeamLockerRoom,
			HomeTeamShotsOnGoal: game.HomeTeamShotsOnGoal,
			HomeTeamScore:       game.HomeTeamScore,
			AwayTeam:            game.AwayTeam,
			AwayTeamID:          game.AwayTeamID,
			AwayTeamRoster:      game.AwayTeamRoster,
			AwayTeamRosterID:    game.AwayTeamRosterID,
			AwayTeamLockerRoom:  game.AwayTeamLockerRoom,
			AwayTeamShotsOnGoal: game.AwayTeamShotsOnGoal,
			AwayTeamScore:       game.AwayTeamScore,
			ScoreKeeper:         game.ScoreKeeper,
			ScoreKeeperID:       game.ScoreKeeperID,
			PrimaryReferee:      game.PrimaryReferee,
			PrimaryRefereeID:    game.PrimaryRefereeID,
			SecondaryReferee:    game.SecondaryReferee,
			SecondaryRefereeID:  game.SecondaryRefereeID,
		})
	}
	return gameModels
}

func MapGameModelToGameStruct(games []models.Game) []structures.Game {
	var gameStructs []structures.Game
	for _, game := range games {
		gameStructs = append(gameStructs, structures.Game{
			Game:      game,
			Optimized: false,
		})
	}
	return gameStructs
}
