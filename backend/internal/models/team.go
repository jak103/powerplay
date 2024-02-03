package models

type Team struct {
	DbModel
	Name     string //`json:"name"`
	LogoPath string //`json:"logo_path"`
	Color    string //`json:"color"`
	Active   bool   //`json:"active"`
	LeagueID uint   //`json:"league_id"`
	League   League `json:"league"`
	// Captain   User   `json:"primary_captain" gorm:"foreignKey:CaptainID"`
	// CaptainID uint   `json:"primary_captain_id"`
}
