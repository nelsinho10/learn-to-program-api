package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/nelsinho10/learn-to-program-api/controllers"
)

// UserRouter agrupador de las rutas correspondientes a usuarios
func UsersRouter(r chi.Router) {
	r.Get("/", controllers.GetAllUsers)
	r.Post("/add", controllers.NewUser)
}