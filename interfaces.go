package config

import (
	"time"

	"github.com/spf13/viper"
)

type M = map[string]any

type Config interface {
	// Load load config from loader.
	Load(loader Loader) error
	// Apply apply config to application.
	Use(hooks ...func(Config))
	// Env get config from env.
	Env(envName string, defaultValue ...any) any
	// Add config to application.
	Add(name string, configuration any)
	// Get config from application.
	Get(path string, defaultValue ...any) any
	// GetString get string type config from application.
	GetString(path string, defaultValue ...string) string
	// GetInt get int type config from application.
	GetInt(path string, defaultValue ...int) int
	// GetBool get bool type config from application.
	GetBool(path string, defaultValue ...bool) bool
	// GetDuration get time.Duration type config from application.
	GetDuration(path string, defaultValue ...time.Duration) time.Duration
}

type Loader interface {
	Load(app *viper.Viper) error
}
