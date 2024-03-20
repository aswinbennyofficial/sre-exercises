package database

import (
    "github.com/golang-migrate/migrate/v4"
    "github.com/golang-migrate/migrate/v4/database/postgres"
    _ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/rs/zerolog/log"
 
)

// MigrateDB performs database migrations using pgx.Conn
func MigrateDB() {
    
	

	driver, err := postgres.WithInstance(SQLDB, &postgres.Config{})
	if err != nil {
		log.Fatal().Err(err).Msg("Error while creating the driver")
	}
    m, err := migrate.NewWithDatabaseInstance(
        "file://./internals/database/migration",
        "postgres", driver)
    if err!=nil{
		log.Fatal().Err(err).Msg("Error while creating the migration instance")
	}
	m.Up()
	log.Info().Msg("Database migrated successfully")

}

