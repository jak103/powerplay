package models

type League struct {
	DbModel
	Name string `json:"name"`
	// Teams []*Team `json:"teams" gorm:"many2many:league_teams"`
}
