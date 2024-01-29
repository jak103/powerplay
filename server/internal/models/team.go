package models

type Team struct {
	dbModel
	Name     string    `json:"name"`
	Captains []User    `json:"captains"`
	Players  []User    `json:"players"`
	LogoPath string    `json:"logo_path"`
	Color    string    `json:"color"`
	Schedule []*Game   `json:"schedule"`
	Stats    TeamStats `json:"stats"`
}
