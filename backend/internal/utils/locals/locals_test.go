package locals

import (
	"testing"

	"github.com/jak103/powerplay/internal/models"
	"github.com/jak103/powerplay/internal/server/services/auth"
	"github.com/jak103/powerplay/internal/utils/unittesting"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRequestingUser(t *testing.T) {
	c := unittesting.FiberCtx()
	user := models.User{
		DbModel: models.DbModel{
			ID: 27,
		},
		Roles: auth.Authenticated,
	}

	SetRequestingUser(c, user)

	r := RequestingUser(c)

	require.NotNil(t, r)
	assert.Equal(t, user.ID, r.ID)
	assert.Equal(t, user.Roles, r.Roles)
}

func TestNilRequestingUser(t *testing.T) {
	c := unittesting.FiberCtx()
	r := RequestingUser(c)

	assert.Nil(t, r)
}
