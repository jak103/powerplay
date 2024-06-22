package models

import "github.com/lib/pq"

type Venue struct {
	DbModel
	Name        string         `json:"name" validate:"required"`
	Address     string         `json:"address" validate:"required"`
	LockerRooms pq.StringArray `json:"locker_rooms" gorm:"type:text[]" validate:"required"`
}
