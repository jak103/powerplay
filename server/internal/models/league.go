package models

type League struct {
	dbModel
	Name     string `json:"name"`
	Teams    []Team `json:"teams"`
	Schedule []Game `json:"schedule"`
}
