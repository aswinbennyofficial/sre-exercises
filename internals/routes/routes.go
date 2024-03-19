package routes

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

// LoadRoutes is a function that loads all the routes for the application
func LoadRoutes(r *chi.Mux) {
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})
	
}