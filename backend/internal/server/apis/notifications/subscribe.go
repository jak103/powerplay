package notifications

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jak103/powerplay/internal/models"
	"github.com/jak103/powerplay/internal/server/apis"
	"github.com/jak103/powerplay/internal/server/services/auth"
	"github.com/jak103/powerplay/internal/server/services/notifications"
	"github.com/jak103/powerplay/internal/utils/locals"
	"github.com/jak103/powerplay/internal/utils/responder"
)

func init() {
	apis.RegisterHandler(fiber.MethodPost, "/notifications/subscribe", auth.Public, subscriptionHandler)
	apis.RegisterHandler(fiber.MethodGet, "/notifications/send", auth.Public, pushNotification)
}

func subscriptionHandler(c *fiber.Ctx) error {
	log := locals.Logger(c)

	log.Info("Handling new subscription")

	// Parse topics
	subscriptionRequest := &models.NotificationSubscription{}
	err := c.BodyParser(subscriptionRequest)
	if err != nil {
		log.WithErr(err).Alert("Failed to parse subscription request")
		return responder.InternalServerError(c)
	}

	// db := db.GetSession(c)
	// err = db.SaveSubscription(subscriptionRequest)
	notifications.SaveSubscription(*subscriptionRequest)
	if err != nil {
		log.WithErr(err).Alert("Failed to save subscription request")
		return responder.InternalServerError(c)
	}

	return responder.Ok(c)
}

func pushNotification(c *fiber.Ctx) error {
	log := locals.Logger(c)
	log.Info("Sending push notification")

	err := notifications.SendNotification(c, "test notifications", "test message")
	if err != nil {
		log.WithErr(err).Error("Failed to send notification")
		return responder.InternalServerError(c)
	}

	return responder.Ok(c)
}
