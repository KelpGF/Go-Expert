package user

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func defaultUser() *User {
	return NewUser("test", "test@g.c", "test")
}

func TestNewUser(t *testing.T) {
	user := defaultUser()

	assert.NotNil(t, user)
	assert.NotEmpty(t, user.ID)
	assert.NotEmpty(t, user.Password)
	assert.Equal(t, "test", user.Name)
	assert.Equal(t, "test@g.c", user.Email)
}

func TestUserComparePassword(t *testing.T) {
	user := defaultUser()

	assert.True(t, user.ComparePassword("test"))
	assert.False(t, user.ComparePassword("test1"))
	assert.NotEqual(t, "test", user.Password)
}
