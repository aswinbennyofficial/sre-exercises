package database

import (
    "github.com/golang-migrate/migrate/v4"
    "github.com/golang-migrate/migrate/v4/database/postgres"
    _ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/rs/zerolog/log"
 
)

// MigrateDB performs database migrations using go-migrate package.
// Migration scripts are saved in /internals/database/migration.
func MigrateDB() {
	// Create a new instance of the postgres driver
	driver, err := postgres.WithInstance(SQLDB, &postgres.Config{})
	if err != nil {
		log.Fatal().Err(err).Caller().Msg("Error while creating the driver")
	}
	// Create a new instance of the migration
    m, err := migrate.NewWithDatabaseInstance(
        "file://./internals/database/migration",
        "postgres", driver)
    if err!=nil{
		log.Fatal().Err(err).Caller().Msg("Error while creating the migration instance")
	}
	// Migrate the database up
	m.Up()

	log.Info().Msg("Database migrated successfully")

}

