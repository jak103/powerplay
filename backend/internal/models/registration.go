package models

type Registration struct {
	DbModel
	SeasonID  uint `json:"season_id"`
	UserID    uint
	User      User       `json:"user"`
	Questions []Question `type:"questions"`
}

type Question struct {
	DbModel
	RegistrationID uint
	Text           string `json:"text"`
	Answer         string `json:"answer"`
	Render         string `json:"render"`
}
