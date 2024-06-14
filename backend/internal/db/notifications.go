package db

import "github.com/jak103/powerplay/internal/models"

func (s session) SaveSubscription(request *models.NotificationSubscription) error {
	result := s.connection.Create(request)
	return result.Error
}

func (s session) GetSubscriptionsByTopic(topic string) ([]models.NotificationSubscription, error) {
	subs := make([]models.NotificationSubscription, 0)
	result := s.connection.Where("topic = ?", topic).First(&subs)

	return resultsOrError(subs, result)

}
