package models

import "github.com/nelsinho10/learn-to-program-api/database"

type User struct {
	Uid      string   `json:"uid,omitempty"`
	Name     string   `json:"name,omitempty"`
	Email    string   `json:"email,omitempty"`
	Password string   `json:"password,omitempty"`
	DType    []string `json:"dgraph.type,omitempty"`
}

// AddUser agregar un usuario a la base de datos
func AddUser(name string, email string, password string) {
	// Definir un usuario
	user := User{
		Name:     name,
		Email:    email,
		Password: password,
		DType:    []string{"User"},
	}
	// Realizar una mutacion a la base de datos
	database.MakeMutation(user)
}

// AllUsers obtener todos los usuario de la base de datos
func AllUsers() []byte {
	// Query para obtener todos los usuarios
	const q = `
		{
			users(func: type(User)) {
				uid
				name
				email
				password
			  }
		}
	`
	// Realizar la consulta a la base de datos
	res := database.MakeQuery(q)
	return res.Json
}
