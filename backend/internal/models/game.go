package models

import "time"

type Game struct {
	dbModel
	Teams           []*Team   `json:"teams" gorm:"many2many:games_teams;"`
	Team1LockerRoom string    `json:"home_locker_room"`
	Team2LockerRoom string    `json:"away_locker_room"`
	Start           time.Time `json:"start"`
	End             time.Time `json:"end"`
	Location        Venue     `json:"location"`
	Staff           Staff     `json:"staff"`

	// TODO Figure out scorekeeping and game stats, I think score is just one of the game stats
	// HomeScore int `json:"home_score"`
	// AwayScore int `json:"away_score"`//
}
