package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/nelsinho10/learn-to-program-api/controllers"
)

// ProgramRouter returns a chi router for programs
func ProgramsRouter(r chi.Router) {
	r.Get("/", controllers.GetPrograms)
	r.Get("/{id}", controllers.GetProgram)
	r.Post("/add", controllers.NewProgram)
	r.Post("/execute/{name}", controllers.ExecuteProgram)
}
