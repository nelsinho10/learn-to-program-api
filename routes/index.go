package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

// Router returns a chi router for the API
func Router() *chi.Mux {
	// Create a new router
	r := chi.NewRouter()

	// CORS setup for all routes
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	// Routes for api
	r.Route("/api", func(r chi.Router) {
		r.Route("/users", UsersRouter)
		r.Route("/programs", ProgramsRouter)
	})

	return r
}
