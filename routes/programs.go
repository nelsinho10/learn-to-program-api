package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/nelsinho10/learn-to-program-api/handlers"
)

// ProgramRouter returns a chi router for programs
func ProgramsRouter(r chi.Router) {
	r.Get("/{offset}/{first}", handlers.GetPrograms)
	r.Get("/{id}", handlers.GetProgram)
	r.Post("/{name}", handlers.NewProgram)
	r.Post("/{name}/execute-program", handlers.ExecuteProgram)
	r.Put("/{id}", handlers.UpdateProgram)
	r.Get("/count", handlers.GetNumberOfPrograms)
	r.Delete("/{id}", handlers.DeleteProgram)
}
