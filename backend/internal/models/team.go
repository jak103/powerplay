package models

type Team struct {
	DbModel
	CorrelationId string `json:"correlation_id"`
	Name          string `json:"name"`
	LogoId        string `json:"logo_id"`
	Color         string `json:"color"`
	LeagueID      uint   `json:"league_id"`
	League        League `json:"league"`
	Roster        Roster `json:"roster"`

	Wins   int `json:"wins"`
	Losses int `json:"losses"`
}
