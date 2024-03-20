package models

// Config is a struct that stores the configuration for the application
type Config struct {
	// The port on which the server should run
	Port string
	JWTSecret string
	PostgresURI string
	LogLevel string
	LogFile string 

	
}