package config

import (

    "os"

    "github.com/rs/zerolog"
    "github.com/rs/zerolog/log"
)

var (
    // appLogFile is a file pointer to the local log file
    appLogFile *os.File
    // appLogger is a logger instance for all log levels
    appLogger  zerolog.Logger
)

// LoadLogger is a function that loads the logger configurations.
// Uses zerolog to do structured logging.
// LoadLogger logs to both stdout and local file and also sets the minimum logLevel
func LoadLogger() {
    // Sets the minimum log level to log
    // remove log level "debug" to disable debug logs
    log.Info().Msg("Log level is "+Configs.LogLevel)

    // Set the log level
    if Configs.LogLevel == "debug" {
        // Logs debugLevel and anything above
        zerolog.SetGlobalLevel(zerolog.DebugLevel)
    }else{
        // Logs Infolevel and anything above it
        zerolog.SetGlobalLevel(zerolog.InfoLevel)
    }
   
    // Opening a local file to save the logs
    AppLogFile, err := os.OpenFile(Configs.LogFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        log.Fatal().Err(err).Caller().Msg("Failed to open app log file")
    }

    // Create multi-level writer to write to both file and stdout
    multiLevelWriter := zerolog.MultiLevelWriter(AppLogFile, os.Stdout)

    // Create logger instance for all log levels
    appLogger = zerolog.New(multiLevelWriter).With().Timestamp().Logger()

    // Set the configured logger as the default logger
    log.Logger = appLogger
}

// CloseLogFiles is a function that closes the opened local log file for writing logs
func CloseLogFiles() {
    // Close log file
    if appLogFile != nil {
        if err := appLogFile.Close(); err != nil {
            log.Error().Err(err).Msg("Failed to close app log file")
        }
    }
}
