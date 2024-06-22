package models

type League struct {
	DbModel
	CorrelationId string `json:"correlation_id"  validate:"required"`
	SeasonID      uint   `json:"season_id" validate:"required"`
	Name          string `json:"name" validate:"required"`
	Teams         []Team `json:"teams"  validate:"required"`
}
