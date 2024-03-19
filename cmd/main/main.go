package main

import (
	"net/http"

	"github.com/aswinbennyofficial/sre-exercises/internals/config"
	"github.com/aswinbennyofficial/sre-exercises/internals/routes"
	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"
)
func init(){
	// Load the config
	err:=config.LoadConfig()
	if err!=nil{
		log.Panic().Err(err).Msg("Error while loading the config")
	}
	log.Info().Msg("Config loaded successfully")
	log.Debug().Msg("Config: Port: "+config.Configs.Port)


	// Load the logger
	config.LoadLogger()
}

func main(){
	// Close the log files
	defer config.CloseLogFiles()

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