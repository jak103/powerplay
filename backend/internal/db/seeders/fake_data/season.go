package fake_data

import (
	"github.com/go-faker/faker/v4"
	"github.com/jak103/powerplay/internal/models"
	"github.com/jak103/powerplay/internal/utils/formatters"
	"gorm.io/gorm"
	"time"
)

type SeasonSeeder struct{}

func (s SeasonSeeder) Seed(db *gorm.DB, args ...interface{}) (interface{}, error) {
	season := models.Season{
		Name:          formatters.CapitalizeFirstLetter(faker.Word()),
		Start:         time.Now(),
		End:           time.Now().AddDate(0, 3, 0),
		Registrations: []models.Registration{},
		Schedule:      []models.Game{},
		LeagueRecords: []models.LeagueRecord{},
	}
	if err := db.FirstOrCreate(&season, season).Error; err != nil {
		return nil, err
	}
	return season, nil
}
