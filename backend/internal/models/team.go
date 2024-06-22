package models

type Team struct {
	DbModel
	CorrelationId string  `json:"correlation_id" validate:"required"`
	Name          string  `json:"name" validate:"required"`
	LogoId        string  `json:"logo_id" validate:"required"`
	Color         string  `json:"color" validate:"required"`
	LeagueID      uint    `json:"league_id" validate:"required"`
	League        *League  `json:"league"`
	Roster        *Roster `json:"roster"`
	RosterID      *uint   `json:"roster_id" validate:"required"`

	Wins   int `json:"wins" validate:"required"`
	Losses int `json:"losses" validate:"required"`
}
