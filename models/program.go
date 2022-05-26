package models

import (
	"io"
	"log"
	"time"

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
func AddProgram(r io.ReadCloser, name string) {

	// Program data
	b, error := io.ReadAll(r)

	if error != nil {
		log.Fatal("Error reading program data", error)
	}

	program := string(b)
	currentTime := time.Now().Format("2006-01-02")

	pd := ProgramData{
		Name:        name,
		Program:     program,
		DateCreated: currentTime,
		DateUpdated: currentTime,
		DType:       []string{"Program"},
	}

	database.MakeMutation(pd)
}

// UpdateProgram update program by id from dgraph
func UpdateProgram(id string, r io.ReadCloser) {
	// Program data
	b, error := io.ReadAll(r)

	if error != nil {
		log.Fatal("Error reading program data", error)
	}

	program := string(b)
	currentTime := time.Now().Format("2006-01-02")

	pd := ProgramData{
		Uid:         id,
		Program:     program,
		DateUpdated: currentTime,
		DType:       []string{"Program"},
	}

	database.MakeMutation(pd)
}

// DeleteProgram delete program by id from dgraph
func DeleteProgram(id string) {
	// database.MakeDelete(id)
}

// AllPrograms returns all programs from dgraph
func AllPrograms(initial string, final string) []byte {

	// Query
	q := `
	{
		programs(func: type(Program),offset: ` + initial + `, first: ` + final + `) {
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

// CountPrograms return number of programs from dgraph
func CountPrograms() []byte {
	// Query
	q := `
	{
		programs_counts(func: type(Program)){
			count(uid)
  		}
	}
	`
	//
	res := database.MakeQuery(q)
	return res.Json
}
