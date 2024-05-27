package models

type League struct {
	DbModel
	CorrelationId string `json:"correlation_id"`
	Name          string `json:"name"`
	Teams         []Team `json:"teams"`
}
