package models

import (
	"github.com/SherClockHolmes/webpush-go"
)

type Topic string

const (
	RSVP         Topic = "new_rsvp"
	CHAT         Topic = "new_chat"
	GAME_UPDATE  Topic = "game_update"
	EVENT_UPDATE Topic = "event_update"
)

// All the JSON is blanked because this is sensitive information and should never go to the front end
type NotificationSubscription struct {
	DbModel
	Topics []Topic `json:"topics"`
	webpush.Subscription
}
