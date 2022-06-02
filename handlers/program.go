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

	uid, error := models.AddProgram(r.Body, name)

	if error != nil {
		http.Error(w, error.Error(), http.StatusInternalServerError)
		return
	}

	// Response with json
	w.Write([]byte(uid))
}

// GetPrograms returns all programs
func GetPrograms(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// Get range page
	offset := chi.URLParam(r, "offset")
	first := chi.URLParam(r, "first")

	b, error := models.AllPrograms(offset, first)

	if error != nil {
		http.Error(w, error.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(b)
}

// GetProgram return program by id
func GetProgram(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := chi.URLParam(r, "id")
	b, error := models.GetProgram(id)

	if error != nil {
		http.Error(w, error.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(b)
}

// ExecuteProgram execute program by name
func ExecuteProgram(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	name := chi.URLParam(r, "id")
	namePyhonFile := fmt.Sprintf("%s.py", name)

	b, error := io.ReadAll(r.Body)

	if error != nil {
		http.Error(w, error.Error(), http.StatusInternalServerError)
		return
	}

	// Create a new file
	error = helpers.CreateFile(namePyhonFile, string(b))

	if error != nil {
		http.Error(w, error.Error(), http.StatusInternalServerError)
		return
	}

	// Run program
	out, error := helpers.RunPythonFile(namePyhonFile)

	if error != nil {
		http.Error(w, error.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(out)
	helpers.DeleteFile(namePyhonFile)
}

// UpdateProgram update program by id
func UpdateProgram(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := chi.URLParam(r, "id")
	_, error := models.UpdateProgram(id, r.Body)

	if error != nil {
		http.Error(w, error.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Program updated"))
}

// GetNumberOfPrograms return number of programs
func GetNumberOfPrograms(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	b, error := models.CountPrograms()

	if error != nil {
		http.Error(w, error.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(b)
}

// DeleteProgram delete program by id
func DeleteProgram(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := chi.URLParam(r, "id")
	error := models.DeleteProgram(id)

	if error != nil {
		http.Error(w, error.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Program deleted"))
}
