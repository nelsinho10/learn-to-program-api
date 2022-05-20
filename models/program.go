package models

import (
	"io"

	"github.com/nelsinho10/learn-to-program-api/database"
)

type ProgramData struct {
	Uid         string   `json:"uid,omitempty"`
	Name        string   `json:"name,omitempty"`
	DateCreated string   `json:"date_created,omitempty"`
	DateUpdated string   `json:"date_updated,omitempty"`
	Program     string   `json:"program"`
	DType       []string `json:"dgraph.type,omitempty"`
}

// AddProgram add new program to dgraph
func AddProgram(r io.ReadCloser) {

	// Program data
	b, error := io.ReadAll(r)

	if error != nil {
		return
	}

	program := string(b)

	pd := ProgramData{
		Name:        "program1",
		Program:     program,
		DateCreated: "2020-01-01",
		DateUpdated: "2020-01-01",
		DType:       []string{"Program"},
	}

	database.MakeMutation(pd)
}

// AllPrograms returns all programs from dgraph
func AllPrograms() []byte {

	// Query
	q := `
	{
		programs(func: type(Program)) {
			uid
			name
			date_created
			date_updated
			program
		  }
	}
	`
	//
	res := database.MakeQuery(q)
	return res.Json
}

// GetProgram return program by id from dgraph
func GetProgram(id string) []byte {
	// Query
	q := `
	{
		program(func: uid(` + id + `)) {
			uid
			name
			date_created
			date_updated
			program
		  }
	}
	`
	//
	res := database.MakeQuery(q)
	return res.Json
}
