package cfg_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ownfitness/template-go/pkg/cfg"
)

func TestNew(t *testing.T) {
	env := map[string]string{
		"PORT":    "3000",
		"DEBUG":   "true",
		"PROJECT": "test-project",
	}

	setup(env)
	defer teardown(env)

	c, err := cfg.New()

	assert.NoError(t, err)
	assert.Equal(t, env["PORT"], c.Port)
	assert.Equal(t, env["PROJECT"], c.Project)
	assert.True(t, c.Debug)

}

func TestNewDefaults(t *testing.T) {
	env := map[string]string{
		"PROJECT": "test-project",
	}

	setup(env)
	defer teardown(env)

	c, err := cfg.New()

	assert.NoError(t, err)
	assert.Equal(t, "8080", c.Port)
	assert.False(t, c.Debug)
}

func TestNewError(t *testing.T) {
	_, err := cfg.New()

	assert.Error(t, err)
}

func setup(e map[string]string) {
	for k, v := range e {
		os.Setenv(k, v)
	}
}

func teardown(e map[string]string) {
	for k, _ := range e {
		os.Unsetenv(k)
	}
}
