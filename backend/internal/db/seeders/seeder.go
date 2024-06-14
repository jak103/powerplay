package seeders

import "gorm.io/gorm"

type Seeder interface {
	Seed(db *gorm.DB, args ...interface{}) (interface{}, error)
}
