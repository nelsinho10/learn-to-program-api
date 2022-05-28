package handlers

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
	// Get name from url
	name := chi.URLParam(r, "name")

	uid := models.AddProgram(r.Body, name)

	// Response with json
	w.Write([]byte(uid))
}

// GetPrograms returns all programs
func GetPrograms(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// Get range page
	offset := chi.URLParam(r, "offset")
	first := chi.URLParam(r, "first")

	b := models.AllPrograms(offset, first)
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
		helpers.Error(w, error)
		return
	}

	// Create a new file
	helpers.CreateFile(namePyhonFile, string(b))

	// Run program
	out := helpers.RunPythonFile(namePyhonFile)
	w.Write(out)
	helpers.DeleteFile(namePyhonFile)
}

// UpdateProgram update program by id
func UpdateProgram(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := chi.URLParam(r, "id")
	models.UpdateProgram(id, r.Body)
	w.Write([]byte("Program updated"))
}

// GetNumberOfPrograms return number of programs
func GetNumberOfPrograms(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	b := models.CountPrograms()
	w.Write(b)
}

// DeleteProgram delete program by id
func DeleteProgram(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := chi.URLParam(r, "id")
	models.DeleteProgram(id)
	w.Write([]byte("Program deleted"))
}
