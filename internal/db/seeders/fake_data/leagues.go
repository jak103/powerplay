package fake_data

import (
	"errors"
	"github.com/go-faker/faker/v4"
	"github.com/jak103/powerplay/internal/models"
	"github.com/jak103/powerplay/internal/utils/unittesting"
	"gorm.io/gorm"
)

type LeagueSeeder struct{}

func (s LeagueSeeder) Seed(db *gorm.DB, args ...interface{}) (interface{}, error) {
	// Expecting the first argument to be the SeasonID
	if len(args) < 1 {
		return nil, errors.New("missing required arguments")
	}
	seasonID, ok := args[0].(uint)
	if !ok {
		return nil, errors.New("invalid type for SeasonID")
	}

	existingNames := make(map[string]bool)
	var createdLeagues []models.League
	for i := 0; i < 4; i++ {
		league := models.League{
			Name:          unittesting.GenerateUniqueName(existingNames),
			CorrelationId: faker.UUIDHyphenated(),
			SeasonID:      seasonID,
			Teams:         []models.Team{},
		}
		if err := db.FirstOrCreate(&league, models.League{
			Name:          league.Name,
			CorrelationId: league.CorrelationId,
			SeasonID:      league.SeasonID,
		}).Error; err != nil {
			return nil, err
		}
		createdLeagues = append(createdLeagues, league)
	}

	return createdLeagues, nil
}
