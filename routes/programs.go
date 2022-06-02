package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/nelsinho10/learn-to-program-api/handlers"
)

// ProgramRouter returns a chi router for programs
func ProgramsRouter(r chi.Router) {
	r.Get("/{offset}/{first}", handlers.GetPrograms)
	r.Get("/{id}", handlers.GetProgram)
	r.Get("/count", handlers.GetNumberOfPrograms)
	r.Post("/{name}", handlers.NewProgram)
	r.Post("/{id}/execute-program", handlers.ExecuteProgram)
	r.Patch("/{id}", handlers.UpdateProgram)
	r.Delete("/{id}", handlers.DeleteProgram)
}
