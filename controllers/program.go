package controllers

import (
	"fmt"
	"io"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/nelsinho10/learn-to-program-api/helpers"
	"github.com/nelsinho10/learn-to-program-api/models"
)

// NewProgram add new program
func NewProgram(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	models.AddProgram(r.Body)
}

// GetPrograms returns all programs
func GetPrograms(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	b := models.AllPrograms()
	w.Write(b)
}

// GetProgram return program by id
func GetProgram(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := chi.URLParam(r, "id")
	b := models.GetProgram(id)
	w.Write(b)
}

// ExecuteProgram execute program by name
func ExecuteProgram(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	name := chi.URLParam(r, "name")
	namePyhonFile := fmt.Sprintf("%s.py", name)

	b, error := io.ReadAll(r.Body)

	if error != nil {
		return
	}

	// Create a new file
	helpers.CreateFile(namePyhonFile, string(b))

	// Run program
	out := helpers.RunPythonFile(namePyhonFile)
	w.Write(out)
}
