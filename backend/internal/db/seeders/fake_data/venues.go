package fake_data

import (
	"github.com/go-faker/faker/v4"
	"github.com/jak103/powerplay/internal/models"
	"github.com/jak103/powerplay/internal/utils/formatters"
	"github.com/jak103/powerplay/internal/utils/unittesting"
	fakeraddress "github.com/jaswdr/faker/v2"
	"gorm.io/gorm"
	"math/rand"
)

type VenueSeeder struct{}

func generateLockerRooms() []string {
	numRooms := rand.Intn(5) + 2
	lockerRooms := make([]string, numRooms)
	for i := range lockerRooms {
		lockerRooms[i] = formatters.CapitalizeFirstLetter(faker.Word())
	}
	return lockerRooms
}

func (s VenueSeeder) Seed(db *gorm.DB, args ...interface{}) (interface{}, error) {
	fakeaddress := fakeraddress.New()
	existingNames := make(map[string]bool)
	var createdVenues []models.Venue
	for i := 0; i < 4; i++ {
		venue := models.Venue{
			Name:        unittesting.GenerateUniqueName(existingNames),
			Address:     fakeaddress.Address().Address(),
			LockerRooms: generateLockerRooms(),
		}
		if err := db.FirstOrCreate(&venue, models.Venue{
			Name:    venue.Name,
			Address: venue.Address,
		}).Error; err != nil {
			return nil, err
		}
		createdVenues = append(createdVenues, venue)
	}

	return createdVenues, nil
}
