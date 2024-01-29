package models

type Registration struct {
	dbModel
	User      User       `json:"user"`
	Questions []Question `type:"questions"`
}

type Question struct {
	dbModel
	Text   string `json:"text"`
	Answer string `json:"answer"`
	Render string `json:"render"`
}
