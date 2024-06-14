package db

import "github.com/jak103/powerplay/internal/models"

func (s Session) SaveLogo(logo *models.Logo) error {
	result := s.Connection.Create(logo)
	return result.Error
}

func (s Session) GetLogoByID(id string) (*models.Logo, error) {
	var logo models.Logo
	err := s.Connection.First(&logo, "id = ?", id)
	return resultOrError(&logo, err)
}
