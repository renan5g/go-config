package config

import (
	"time"

	"github.com/spf13/cast"
	"github.com/spf13/viper"
)

type Application struct {
	vip *viper.Viper
}

func NewConfig() *Application {
	app := &Application{}
	app.vip = viper.New()
	app.vip.AutomaticEnv()
	return app
}

func (app *Application) Instance() *viper.Viper {
	return app.vip
}

func (app *Application) Load(loader Loader) error {
	return loader.Load(app.vip)
}

func (app *Application) Use(hooks ...func(Config)) {
	for _, hook := range hooks {
		hook(app)
	}
}

// Env Get config from env.
func (app *Application) Env(envName string, defaultValue ...any) any {
	value := app.Get(envName, defaultValue...)
	if cast.ToString(value) == "" {
		return Default(defaultValue...)
	}

	return value
}

// Add config to application.
func (app *Application) Add(name string, configuration any) {
	app.vip.Set(name, configuration)
}

// Get config from application.
func (app *Application) Get(path string, defaultValue ...any) any {
	if !app.vip.IsSet(path) {
		return Default(defaultValue...)
	}
	return app.vip.Get(path)
}

// GetString get string type config from application.
func (app *Application) GetString(path string, defaultValue ...string) string {
	if !app.vip.IsSet(path) {
		return Default(defaultValue...)
	}
	return app.vip.GetString(path)
}

// GetInt get int type config from application.
func (app *Application) GetInt(path string, defaultValue ...int) int {
	if !app.vip.IsSet(path) {
		return Default(defaultValue...)
	}
	return app.vip.GetInt(path)
}

// GetBool get bool type config from application.
func (app *Application) GetBool(path string, defaultValue ...bool) bool {
	if !app.vip.IsSet(path) {
		return Default(defaultValue...)
	}
	return app.vip.GetBool(path)
}

// GetDuration get time.Duration type config from application
func (app *Application) GetDuration(path string, defaultValue ...time.Duration) time.Duration {
	if !app.vip.IsSet(path) {
		return Default(defaultValue...)
	}
	return app.vip.GetDuration(path)
}
