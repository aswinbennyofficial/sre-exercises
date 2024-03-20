package database

import (
	"context"
	"database/sql"
	"github.com/aswinbennyofficial/sre-exercises/internals/config"
	"github.com/jackc/pgx/v5"
	"github.com/rs/zerolog/log"
)

var (
	SQLDB  *sql.DB
	PgxConn *pgx.Conn
)

// ConnectDB is a function that connects to the database
func ConnectDB() {
	// Connect to PostgreSQL using pgx
	pgxConfig, err := pgx.ParseConfig(config.Configs.PostgresURI)
	if err != nil {
		log.Fatal().Err(err).Msg("Error parsing PostgreSQL config")
	}

	conn, err := pgx.ConnectConfig(context.Background(), pgxConfig)
	if err != nil {
		log.Fatal().Err(err).Msg("Error connecting to PostgreSQL database")
	}
	PgxConn = conn

	// Open a connection to the MySQL database using database/sql
	sqlConn, err := sql.Open("postgres", config.Configs.PostgresURI)
	if err != nil {
		log.Fatal().Err(err).Msg("Error connecting to SQL database")
	}
	SQLDB = sqlConn

	// Ping PostgreSQL
	err = PgxConn.Ping(context.Background())
	if err != nil {
		log.Fatal().Err(err).Msg("Error while pinging PostgreSQL database")
	}

	log.Info().Msg("Connected to both PostgreSQL and SQL databases")
}

// CloseDB is a function to close database connections
func CloseDB() {
	if SQLDB != nil {
		err := SQLDB.Close()
		if err != nil {
			log.Error().Err(err).Msg("Error closing SQL database connection")
		}
		log.Info().Msg("SQL database connection closed")
	}
	if PgxConn != nil {
		err := PgxConn.Close(context.Background())
		if err != nil {
			log.Error().Err(err).Msg("Error closing PostgreSQL connection")
		}
		log.Info().Msg("PostgreSQL connection closed")
	}
}
