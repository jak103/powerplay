package fake_data

import (
	"github.com/go-faker/faker/v4"
	"github.com/jak103/powerplay/internal/models"
	"gorm.io/gorm"
	"time"
)

type SeasonSeeder struct{}

func (s SeasonSeeder) Seed(db *gorm.DB) error {
	season := models.Season{
		Name:          faker.Word(),
		Start:         time.Now(),
		End:           time.Now().AddDate(0, 3, 0),
		Registrations: []models.Registration{},
		Schedule:      []models.Game{},
		Leagues:       []models.League{},
	}
	if err := db.FirstOrCreate(&season, season).Error; err != nil {
		return err
	}
	return nil
}
