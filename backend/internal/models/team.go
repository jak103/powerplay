package models

type Team struct {
	dbModel
	Name     string `json:"name"`
	Captains []User `json:"captains" gorm:"many2many:team_captains;"`
	Players  []User `json:"players" gorm:"many2many:team_players;"`
	LogoPath string `json:"logo_path"`
	Color    string `json:"color"`
	Active   bool   `json:"active"`

	// Stats    TeamStats `json:"stats"`
}
