package fake_data

import (
	"errors"
	"github.com/go-faker/faker/v4"
	"github.com/jak103/powerplay/internal/models"
	"github.com/jak103/powerplay/internal/utils/unittesting"
	fakercolor "github.com/jaswdr/faker/v2"
	"gorm.io/gorm"
)

type TeamSeeder struct{}

func (s TeamSeeder) Seed(db *gorm.DB, args ...interface{}) (interface{}, error) {
	fake := fakercolor.New()
	// Expecting the first argument to be the LeagueID
	if len(args) < 1 {
		return nil, errors.New("missing required arguments")
	}
	leagueID, ok := args[0].(uint)
	if !ok {
		return nil, errors.New("invalid type for LeagueID")
	}

	existingNames := make(map[string]bool)
	var createdTeams []models.Team
	for i := 0; i < 4; i++ {
		team := models.Team{
			CorrelationId: faker.UUIDHyphenated(),
			Name:          unittesting.GenerateUniqueName(existingNames),
			Color:         fake.Color().Hex(),
			LeagueID:      leagueID,
			Wins:          0,
			Losses:        0,
		}
		if err := db.FirstOrCreate(&team, models.Team{
			CorrelationId: team.CorrelationId,
			Name:          team.Name,
			Color:         team.Color,
			LeagueID:      team.LeagueID,
		}).Error; err != nil {
			return nil, err
		}
		createdTeams = append(createdTeams, team)
	}

	return createdTeams, nil
}
