package database

import (
	"context"
	"database/sql"
	"github.com/aswinbennyofficial/sre-exercises/internals/config"
	// "github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog/log"
)

var (
	// SQLDB is a pointer to the SQL database connection object
	SQLDB *sql.DB
	// PgxConn is a pointer to the PostgreSQL connection object
	PgxPool *pgxpool.Pool
)

// ConnectDB is a function that connects to the postgres database.
// Generates two types of connection object *pgx.Conn as PgxConn and *sql.DB as SQLDB
func ConnectDB() {
	

	// Connect to PostgreSQL using pgx to get pgx.Conn object
	pgxpool, err := pgxpool.New(context.Background(), config.Configs.PostgresURI)
	if err != nil {
		log.Panic().Err(err).Caller().Msg("Error connecting to PostgreSQL database via pgx")
	}
	PgxPool = pgxpool

	// connect to SQL database using database/sql to get *sql.DB object for migrations
	sqlConn, err := sql.Open("postgres", config.Configs.PostgresURI)
	if err != nil {
		log.Panic().Err(err).Caller().Msg("Error connecting to SQL database via database/sql")
	}
	SQLDB = sqlConn

	// Ping PostgreSQL
	err = pgxpool.Ping(context.Background())
	if err != nil {
		log.Panic().Err(err).Msg("Error while pinging PostgreSQL database")
	}

	log.Info().Msg("Connected to PostgreSQL")
}

// CloseDB is a function to close database connections
func CloseDB() {
	if SQLDB != nil {
		err := SQLDB.Close()
		if err != nil {
			log.Error().Err(err).Caller().Msg("Error closing SQL database connection")
		}else{
			log.Info().Msg("SQLDB connection closed")
		}
	}

	if PgxPool != nil {
		PgxPool.Close()
		log.Info().Msg("PgxPool connection closed")
	}
}
