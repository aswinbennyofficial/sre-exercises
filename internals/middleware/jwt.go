package middleware

import (
	// "net/http"

	"github.com/aswinbennyofficial/sre-exercises/internals/config"
	"github.com/rs/zerolog/log"

	"github.com/go-chi/jwtauth/v5"
)

var TokenAuth *jwtauth.JWTAuth

// InitJWT is a function that initializes the JWT middleware.
// It uses the JWTSecret from the config to sign the JWT tokens
func InitJWT() {
	TokenAuth = jwtauth.New("HS256", []byte(config.Configs.JWTSecret), nil)
	log.Info().Msg("JWT initialized")
}


