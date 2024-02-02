package models

type League struct {
	dbModel
	Name    string   `json:"name"`
	Teams   []Team   `json:"teams"`
	Seasons []Season `json:"seasons"`
}
