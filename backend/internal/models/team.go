package models

type Team struct {
	DbModel
	CorrelationId  string       `json:"correlation_id"`
	Name           string       `json:"name"`
	LogoId         string       `json:"logo_id"`
	Color          string       `json:"color"`
	LeagueRecordID uint         `json:"league_record_id"`
	LeagueRecord   LeagueRecord `json:"league_record"`
	Roster         *Roster      `json:"roster"`
	RosterID       *uint        `json:"roster_id"`

	Wins   int `json:"wins"`
	Losses int `json:"losses"`
}
