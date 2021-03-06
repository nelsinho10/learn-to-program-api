package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

// Router returns a chi router for the API
func Router() *chi.Mux {
	// Create a new router
	r := chi.NewRouter()

	// Logger middleware
	r.Use(middleware.Logger)

	// CORS setup for all routes
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	// Routes for api
	r.Route("/api/v1", func(r chi.Router) {
		r.Route("/programs", ProgramsRouter)
	})
	return r
}
