package db

import "github.com/jak103/powerplay/internal/models"

func (s *session) GetGamesBySeason(seasonId int) (*[]models.Game, error) {
    games := make([]models.Game, 0)

    //TODO: research if SeasonId is actually the database field, it probably isn't
	result := s.connection.Where("SeasonId = ?", seasonId).Find(&games)

    return resultOrError(&games, result)
}
