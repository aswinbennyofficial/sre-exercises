package models

type Config struct {
	// The port on which the server should run
	Port string
	JWTSecret string
	PostgresURI string
	LogLevel string 

	
}