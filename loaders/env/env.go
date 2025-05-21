package env

import (
	"github.com/spf13/viper"
)

type EnvLoader struct {
	envPath string
}

func NewEnvLoader(envPath string) *EnvLoader {
	return &EnvLoader{envPath: envPath}
}

func (loader *EnvLoader) Load(app *viper.Viper) error {
	app.AutomaticEnv()
	app.SetConfigType("env")
	app.SetConfigFile(loader.envPath)
	if err := app.ReadInConfig(); err != nil {
		return err
	}
	return nil
}
