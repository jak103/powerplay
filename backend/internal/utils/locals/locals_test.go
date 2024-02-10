package locals

import (
	"testing"

	"github.com/jak103/powerplay/internal/models"
	"github.com/jak103/powerplay/internal/server/services/auth"
	"github.com/jak103/powerplay/internal/utils/unittesting"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestKeyRecord(t *testing.T) {
	c := unittesting.FiberCtx()
	record := models.KeyRecord{
		DbModel: models.DbModel{
			ID: 27,
		},
		UserId: 1,
		Roles:  auth.Authenticated,
	}

	SetKeyRecord(c, record)

	r := KeyRecord(c)

	require.NotNil(t, r)
	assert.Equal(t, record.ID, r.ID)
	assert.Equal(t, record.UserId, r.UserId)
	assert.Equal(t, record.Roles, r.Roles)
}

func TestNilKeyRecord(t *testing.T) {
	c := unittesting.FiberCtx()
	r := KeyRecord(c)

	assert.Nil(t, r)
}
