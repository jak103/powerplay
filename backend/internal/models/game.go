package models

import "time"

type Game struct {
	DbModel
	SeasonID        uint
	Teams           []Team    `json:"teams" gorm:"many2many:games_teams;"`
	Team1LockerRoom string    `json:"home_locker_room"`
	Team2LockerRoom string    `json:"away_locker_room"`
	Start           time.Time `json:"start"`
	End             time.Time `json:"end"`
	Venue           Venue     `json:"venue"`
	VenueID         uint

	// ManagerOnCall    User      `json:"manager_on_call"`
	// ManagerOnCallID  uint
	// ScoreKeeper      User `json:"score_keeper"`
	// PrimaryReferee   User `json:"primary_referee"`
	// SecondaryReferee User `json:"secondary_referee"`

	// TODO Figure out scorekeeping and game stats, I think score is just one of the game stats
	// HomeScore int `json:"home_score"`
	// AwayScore int `json:"away_score"`//
}
