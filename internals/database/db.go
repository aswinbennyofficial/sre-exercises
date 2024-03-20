package database

import (
	"github.com/aswinbennyofficial/sre-exercises/internals/config"
	"github.com/jackc/pgx/v5"
	"github.com/rs/zerolog/log"
	"context"
)

var Conn *pgx.Conn

// ConnectDB is a function that connects to the database
func ConnectDB() {
	// Connect to the database
	conn, err := pgx.Connect(context.Background(), config.Configs.PostgresURI)
	if err != nil {
		log.Fatal().Err(err).Msg("Error while connecting to the database")
	}

	Conn=conn;
	err=Conn.Ping(context.Background())
	if err!=nil{
		log.Fatal().Err(err).Msg("Error while pinging the database")
	}
}

func CloseDB(){
	Conn.Close(context.Background())
	log.Info().Msg("Database connection closed")
}
