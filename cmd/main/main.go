package main

import (
	"net/http"

	"github.com/aswinbennyofficial/sre-exercises/internals/config"
	"github.com/aswinbennyofficial/sre-exercises/internals/database"
	"github.com/aswinbennyofficial/sre-exercises/internals/routes"
	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"
)



func init(){
	// Load the configs from env and config.yaml
	err:=config.LoadConfig()
	if err!=nil{
		log.Panic().Err(err).Msg("Error while loading the config")
	}
	// Load the logger
	config.LoadLogger()

	// Log the successful loading of the config
	log.Info().Msg("Config loaded successfully")
	log.Debug().Msg("Config: Port: "+config.Configs.Port)
	log.Debug().Msg("Config: JWT Secret: "+config.Configs.JWTSecret)
	log.Debug().Msg("Config: Postgres URI: "+config.Configs.PostgresURI)
	log.Debug().Msg("Config: Log Level: "+config.Configs.LogLevel)
	log.Debug().Msg("Config: Log File: "+config.Configs.LogFile)

	
	// Connect to the database
	database.ConnectDB()
	database.MigrateDB()

	//
	
}

func main(){
	// Close the log files
	defer config.CloseLogFiles()
	// Close the database connection
	defer database.CloseDB()


	// Initialize Chi router
	r := chi.NewRouter()
	// Load the routes
	routes.LoadRoutes(r)

	// Start the server
	log.Info().Msg("Starting the server on port "+config.Configs.Port+"...")
	err:=http.ListenAndServe(":"+config.Configs.Port,r)
	if err!=nil{
		log.Panic().Err(err).Msg("Error while starting the server")
	}
}