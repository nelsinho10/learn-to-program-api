package routes

import "github.com/go-chi/chi/v5"

// Router agrupador de todas las sub rutas de la aplicacion
func Router() *chi.Mux {
	// Inicializando el cliente de la base de datos
	r := chi.NewRouter()

	// Agrupador de las rutas
	r.Route("/api", func(r chi.Router) {
		r.Route("/user", UsersRouter)
	})

	return r
}
