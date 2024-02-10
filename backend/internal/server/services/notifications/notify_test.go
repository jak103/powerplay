package notifications

import (
	"testing"

	"github.com/SherClockHolmes/webpush-go"
	"github.com/jak103/powerplay/internal/utils/log"
	"github.com/stretchr/testify/assert"
)

func TestGenerateVapidKeys(t *testing.T) {
	privateKey, publicKey, err := webpush.GenerateVAPIDKeys()
	assert.Nil(t, err)

	log.Info("Private: %s\nPublic: %s", privateKey, publicKey)
}
