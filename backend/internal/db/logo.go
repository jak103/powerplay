package db

import "github.com/jak103/powerplay/internal/models"

func (s session) SaveLogo(logo *models.Logo) error {
	result := s.connection.Create(logo)
	return result.Error
}

func (s session) GetLogoByID(id string) (*models.Logo, error) {
	var logo models.Logo
	err := s.connection.First(&logo, "id = ?", id)
	return resultOrError(&logo, err)
}
