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
	config.LoadConfig()
	
	// Load the logger
	config.LoadLogger()

	
	// Connect to the database
	database.ConnectDB()
	database.MigrateDB()

	
	
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