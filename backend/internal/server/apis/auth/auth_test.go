package auth

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidatePassword(t *testing.T) {
	assert.True(t, validatePassword("test", "7pDrgB/2lCRP3i/PHkfyvn5Mw7RV++SyX+L2pcyRI5c=", "1234567890"))
}
