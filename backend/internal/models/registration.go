package models

type Registration struct {
	DbModel
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
