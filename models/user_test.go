package models_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ownfitness/template-go/models"
)

func TestUser(t *testing.T) {
	data := models.User{
		Name:  "Tester",
		Email: "test@example.com",
	}

	valid, response := data.Validate()

	assert.Empty(t, response)
	assert.True(t, valid)
}

func TestUserErrorName(t *testing.T) {
	data := models.User{
		Name:  "jo",
		Email: "bad-email",
	}

	valid, response := data.Validate()

	assert.Equal(t, "username less then 3 characters", response)
	assert.False(t, valid)
}

func TestUserErrorEmail(t *testing.T) {
	data := models.User{
		Name:  "Tester",
		Email: "bad-email",
	}

	valid, response := data.Validate()

	assert.Equal(t, "not valid email address", response)
	assert.False(t, valid)
}
