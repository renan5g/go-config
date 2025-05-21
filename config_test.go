package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOsVariables(t *testing.T) {
	assert.Nil(t, os.Setenv("APP_KEY", "12345678901234567890123456789013"))
	assert.Nil(t, os.Setenv("APP_NAME", "app"))
	assert.Nil(t, os.Setenv("APP_PORT", "3306"))
	assert.Nil(t, os.Setenv("APP_DEBUG", "true"))

	config := NewConfig()

	assert.Equal(t, "12345678901234567890123456789013", config.GetString("APP_KEY"))
	assert.Equal(t, "app", config.GetString("APP_NAME"))
	assert.Equal(t, 3306, config.GetInt("APP_PORT"))
	assert.True(t, config.GetBool("APP_DEBUG"))
}

func TestConfig(t *testing.T) {
	config := NewConfig()
	config.Use(func(c Config) {
		c.Add("app", M{
			"name":  "app",
			"port":  3306,
			"debug": true,
		})
	})

	assert.Equal(t, "app", config.GetString("app.name"))
	assert.Equal(t, 3306, config.GetInt("app.port"))
	assert.True(t, config.GetBool("app.debug"))
}
