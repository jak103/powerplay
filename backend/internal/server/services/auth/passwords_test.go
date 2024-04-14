package auth

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHash(t *testing.T) {
	hash1, salt1, err := HashPassword("test_password", "")
	assert.NoError(t, err)

	hash2, salt2, err := HashPassword("test_password", salt1)

	assert.Equal(t, hash1, hash2)
	assert.Equal(t, salt1, salt2)
	assert.NoError(t, err)

	assert.True(t, ComparePassword("test_password", salt1, hash1))
}
