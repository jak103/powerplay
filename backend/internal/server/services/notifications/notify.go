package notifications

import (
	"io"

	"github.com/SherClockHolmes/webpush-go"
	"github.com/gofiber/fiber/v2"
	"github.com/jak103/powerplay/internal/config"
	"github.com/jak103/powerplay/internal/models"
	"github.com/jak103/powerplay/internal/utils/locals"
	"github.com/jak103/powerplay/internal/utils/log"
)

type Topic string

var subs []models.NotificationSubscription

func SaveSubscription(sub models.NotificationSubscription) {
	log.Info("saving sub: %v", sub)
	subs = append(subs, sub)
}

func SendNotification(c *fiber.Ctx, topic string, message string) error {
	log := locals.Logger(c)
	// get all subscriptions to the topic from DB
	// db := db.GetSession(c)
	// subs, err := db.GetSubscriptionsByTopic(topic)
	// if err != nil {
	// 	log.WithErr(err).Alert("Failed to get subscriptions from DB")
	// 	return err
	// }

	// for each subscription
	log.Info("Sending notification to %v subscribers", len(subs))
	for _, sub := range subs {
		// Send Notification
		resp, err := webpush.SendNotification([]byte("Test notification"), &sub.Subscription, &webpush.Options{
			Subscriber:      "https://localhost:9001",
			VAPIDPublicKey:  config.Vars.VapidPublicKey,
			VAPIDPrivateKey: config.Vars.VapidPrivateKey,
			TTL:             30,
		})
		if err != nil {
			log.WithErr(err).Error("Failed to send notification")
			return err
		}
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.WithErr(err).Error("Failed to read response body")
		}
		log.Info("resp.status: %s", resp.Status)
		log.Debug("body: %q", body)
	}

	return nil
}
