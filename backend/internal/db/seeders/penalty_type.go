package seeders

import (
	"github.com/jak103/powerplay/internal/models"
	"gorm.io/gorm"
)

type PenaltyTypeSeeder struct{}

func (pt PenaltyTypeSeeder) Seed(db *gorm.DB, args ...interface{}) (interface{}, error) {
	penaltyTypes := []models.PenaltyType{
		{Name: "Tripping", Duration: 2, Severity: "minor"},
		{Name: "Hooking", Duration: 2, Severity: "minor"},
		{Name: "Slashing", Duration: 2, Severity: "minor"},
		{Name: "Interference", Duration: 2, Severity: "minor"},
		{Name: "Holding", Duration: 2, Severity: "minor"},
		{Name: "High-sticking", Duration: 2, Severity: "minor"},
		{Name: "Fighting", Duration: 5, Severity: "major"},
		{Name: "Charging", Duration: 5, Severity: "major"},
		{Name: "Boarding", Duration: 5, Severity: "major"},
		{Name: "Cross-checking", Duration: 5, Severity: "major"},
		{Name: "Unsportsmanlike Conduct", Duration: 10, Severity: "misconduct"},
		{Name: "Game misconduct", Duration: 0, Severity: "game_misconduct"},
		{Name: "Match Penalty", Duration: 0, Severity: "match"},
	}

	var createdPenaltyTypes []models.PenaltyType
	for _, penaltyType := range penaltyTypes {
		if err := db.FirstOrCreate(&penaltyType, models.PenaltyType{Name: penaltyType.Name, Duration: penaltyType.Duration, Severity: penaltyType.Severity}).Error; err != nil {
			return nil, err
		}
		createdPenaltyTypes = append(createdPenaltyTypes, penaltyType)
	}
	return createdPenaltyTypes, nil
}
