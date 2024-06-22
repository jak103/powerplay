package models

import "time"

type Season struct {
	DbModel
	Name          string         `json:"name" validate:"required"`
	Start         time.Time      `json:"start" validate:"required"`
	End           time.Time      `json:"end" validate:"required"`
	Registrations []Registration `json:"registrations"`
	Schedule      []Game         `json:"schedule"`
	Leagues       []League       `json:"leagues"`
}