package routes

import (
	"net/http"

	"github.com/aswinbennyofficial/sre-exercises/internals/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"github.com/go-chi/httprate"
	"time"
)

// LoadRoutes is a function that loads all the routes for the application
func LoadRoutes(r *chi.Mux) {
	// Load jwt
	middleware.InitJWT()
  
	r.Route("/v1",func(r chi.Router) {
		r.Use(httprate.LimitByIP(150, 1*time.Minute))
	  	r.Use(jwtauth.Verifier(middleware.TokenAuth))
	  	r.Use(jwtauth.Authenticator(middleware.TokenAuth))
  
	 	 // Protected routes
	  	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			
		// Handle authorized request
	  })
	})
  
	// Public routes (optional)
	r.Get("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
	  w.Write([]byte("OK"))
	})
  }
