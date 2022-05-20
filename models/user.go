package models

import "github.com/nelsinho10/learn-to-program-api/database"

type User struct {
	Uid      string   `json:"uid,omitempty"`
	Name     string   `json:"name,omitempty"`
	Email    string   `json:"email,omitempty"`
	Password string   `json:"password,omitempty"`
	DType    []string `json:"dgraph.type,omitempty"`
}

// AddUser add new user to dgraph
func AddUser(name string, email string, password string) {

	// User data
	user := User{
		Name:     name,
		Email:    email,
		Password: password,
		DType:    []string{"User"},
	}
	database.MakeMutation(user)
}

// AllUsers returns all users from dgraph
func AllUsers() []byte {
	// Query
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
	res := database.MakeQuery(q)
	return res.Json
}
