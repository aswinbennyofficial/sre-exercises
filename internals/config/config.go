package config

import (
	"github.com/aswinbennyofficial/sre-exercises/internals/models"
	"github.com/spf13/viper"
)

var Configs models.Config

func LoadConfig()(error) {
	// Set the path to the YAML file
    viper.SetConfigFile("config.yaml")
	err := viper.ReadInConfig()
    if err != nil {
        return err
    }

	Configs.Port = viper.GetString("server.port")
	Configs.JWTSecret = viper.GetString("jwt.secret")
	Configs.PostgresURI = viper.GetString("postgres.uri")
	Configs.LogLevel = viper.GetString("logs.level")

	return nil
}
