package controllers

import (
	"net/http"

	"github.com/nelsinho10/learn-to-program-api/models"
)

// GetAllUsers
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	data := models.AllUsers()
	w.Write(data)
}

// NewUser
func NewUser(w http.ResponseWriter, r *http.Request) {
	models.AddUser("Nelson", "nelson@gmail.com", "123")
	w.Write([]byte("Save Data"))
}
