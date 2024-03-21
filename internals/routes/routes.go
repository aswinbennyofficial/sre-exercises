package routes

import (
	"net/http"

	"github.com/aswinbennyofficial/sre-exercises/internals/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"github.com/go-chi/httprate"
	"time"
)

// LoadRoutes is a function that loads all the http routes for the application
func LoadRoutes(r *chi.Mux) {
	// Initialize JWT middleware to get the TokenAuth object
	middleware.InitJWT()
  


	r.Route("/v1",func(r chi.Router) {
		// Middleware for ratelimiting
		r.Use(httprate.LimitByIP(150, 1*time.Minute))
		// Middleware JWT authorisation
	  	r.Use(jwtauth.Verifier(middleware.TokenAuth))
	  	r.Use(jwtauth.Authenticator(middleware.TokenAuth))
  
	 	 // Protected routes
	  	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			
			r.Header.Set("Content-Type", "application/json")
			w.Write([]byte(`{"message":"Hello, World!"}`))
	
	  })
	})
  


	// Public routes 
	r.Get("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
	  w.Write([]byte("OK"))
	})
  }
