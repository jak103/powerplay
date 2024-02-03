package models

import "time"

type Season struct {
	DbModel
	Name  string    `json:"name"`
	Start time.Time `json:"start"`
	End   time.Time `json:"end"`
	// Registrations []Registration `json:"registrations"`
	Schedule []Game   `json:"schedule"`
	Rosters  []Roster `json:"rosters"`
}
