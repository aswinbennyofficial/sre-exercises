package database

import (
	"context"
	"database/sql"
	"github.com/aswinbennyofficial/sre-exercises/internals/config"
	"github.com/jackc/pgx/v5"
	"github.com/rs/zerolog/log"
)

var (
	// SQLDB is a pointer to the SQL database connection object
	SQLDB  *sql.DB
	// PgxConn is a pointer to the PostgreSQL connection object
	PgxConn *pgx.Conn
)

// ConnectDB is a function that connects to the postgres database.
// Generates two types of connection object *pgx.Conn as PgxConn and *sql.DB as SQLDB 
func ConnectDB() {
	// Connect to PostgreSQL using pgx
	pgxConfig, err := pgx.ParseConfig(config.Configs.PostgresURI)
	if err != nil {
		log.Fatal().Err(err).Msg("Error parsing PostgreSQL config")
	}
	
	// Connect to PostgreSQL using pgx to get pgx.Conn object
	pgxconn, err := pgx.ConnectConfig(context.Background(), pgxConfig)
	if err != nil {
		log.Fatal().Err(err).Caller().Msg("Error connecting to PostgreSQL database via pgx")
	}
	PgxConn = pgxconn



	// connect to SQL database using database/sql to get *sql.DB object for migrations
	sqlConn, err := sql.Open("postgres", config.Configs.PostgresURI)
	if err != nil {
		log.Fatal().Err(err).Caller().Msg("Error connecting to SQL database via database/sql")
	}
	SQLDB = sqlConn



	// Ping PostgreSQL
	err = PgxConn.Ping(context.Background())
	if err != nil {
		log.Fatal().Err(err).Msg("Error while pinging PostgreSQL database")
	}

	
	log.Info().Msg("Connected to PostgreSQL")
}

// CloseDB is a function to close database connections 
func CloseDB() {
	if SQLDB != nil {
		err := SQLDB.Close()
		if err != nil {
			log.Error().Err(err).Caller().Msg("Error closing SQL database connection")
		}
		log.Info().Msg("SQL database connection closed")
	}


	if PgxConn != nil {
		err := PgxConn.Close(context.Background())
		if err != nil {
			log.Error().Err(err).Caller().Msg("Error closing PostgreSQL connection")
		}

		log.Info().Msg("Postgre connection closed")
	}
}
