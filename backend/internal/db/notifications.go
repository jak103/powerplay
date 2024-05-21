package db

import "github.com/jak103/powerplay/internal/models"

func (s Session) SaveSubscription(request *models.NotificationSubscription) error {
	result := s.Connection.Create(request)
	return result.Error
}

func (s Session) GetSubscriptionsByTopic(topic string) ([]models.NotificationSubscription, error) {
	subs := make([]models.NotificationSubscription, 0)
	result := s.Connection.Where("topic = ?", topic).First(&subs)

	return resultsOrError(subs, result)

}
