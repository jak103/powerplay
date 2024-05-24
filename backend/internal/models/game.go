package models

import "time"

type Game struct {
	DbModel
	SeasonID        uint
	Teams           []Team    `json:"teams" gorm:"many2many:games_teams;"`
	Team1LockerRoom string    `json:"home_locker_room"`
	Team2LockerRoom string    `json:"away_locker_room"`
	Start           time.Time `csv:"-"`
	StartDate       string    `csv:"Start_Date"`
	StartTime       string    `csv:"Start_Time"`
	End             time.Time `csv:"-"`
	EndDate         string    `csv:"End_Date"`
	EndTime         string    `csv:"End_Time"`
	Venue           Venue     `json:"venue"`
	VenueID         uint
	IsEarly         bool   `csv:"-"`
	Optimized       bool   `csv:"-"`
	League          string `csv:"-"` // Not in CSV
	IsBye           bool   `csv:"-"`

	// ManagerOnCall    User      `json:"manager_on_call"`
	// ManagerOnCallID  uint
	// ScoreKeeper      User `json:"score_keeper"`
	// PrimaryReferee   User `json:"primary_referee"`
	// SecondaryReferee User `json:"secondary_referee"`

	// TODO Figure out scorekeeping and game stats, I think score is just one of the game stats
	// HomeScore int `json:"home_score"`
	// AwayScore int `json:"away_score"`//
}
