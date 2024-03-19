package main

import (
	"net/http"

	"github.com/aswinbennyofficial/sre-exercises/internals/config"
	"github.com/aswinbennyofficial/sre-exercises/internals/routes"
	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"
)
func init(){
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
	log.Debug().Msg("Starting the server")
	err:=http.ListenAndServe(":8080",r)
	if err!=nil{
		log.Panic().Err(err).Msg("Error while starting the server")
	}
}