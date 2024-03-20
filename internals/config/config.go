package config

import (
	"github.com/aswinbennyofficial/sre-exercises/internals/models"
	"github.com/spf13/viper"
	"github.com/rs/zerolog/log"
)

// Configs is a variable that stores the configuration
var Configs models.Config

// LoadConfig is a function that loads the configuration from the config.yaml and environment variables
// and stores it in the Configs variable.
// It uses viper to read the config file 
func LoadConfig(){
	// Set the path to the YAML file
    viper.SetConfigFile("config.yaml")
	// Read the config file
	err := viper.ReadInConfig()
    if err != nil {
        log.Fatal().Err(err).Caller().Msg("Error reading the config file")
    }

	// Assigning the variable Configs with configurations
	Configs.Port = viper.GetString("server.port")
	Configs.JWTSecret = viper.GetString("jwt.secret")
	Configs.PostgresURI = viper.GetString("postgres.uri")
	Configs.LogLevel = viper.GetString("logs.level")
	Configs.LogFile = viper.GetString("logs.file")

	log.Info().Msg("Configs loaded successfully")

}
