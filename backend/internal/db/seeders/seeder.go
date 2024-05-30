package seeders

import "gorm.io/gorm"

type Seeder interface {
	Seed(db *gorm.DB) error
}
