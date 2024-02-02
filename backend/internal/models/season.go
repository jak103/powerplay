package models

import "time"

type Season struct {
	dbModel
	Name          string         `json:"name"`
	Start         time.Time      `json:"start"`
	End           time.Time      `json:"end"`
	Registrations []Registration `json:"registrations"`
	Schedule      []Game         `json:"schedule"`
}
