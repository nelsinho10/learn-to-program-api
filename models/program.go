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
func AddProgram(r io.ReadCloser, name string) (string, error) {

	// Program data
	b, error := io.ReadAll(r)

	if error != nil {
		log.Println(error)
		return "", error

	}

	program := string(b)
	currentTime := time.Now().Format("2006-01-02  15:04:05")

	pd := ProgramData{
		Name:        name,
		Program:     program,
		DateCreated: currentTime,
		DateUpdated: currentTime,
		DType:       []string{"Program"},
	}

	uid, error := database.MakeMutationAdd(pd)

	if error != nil {
		log.Println(error)
		return "", error
	}

	return uid, nil
}

// UpdateProgram update program by id from dgraph
func UpdateProgram(id string, r io.ReadCloser) (string, error) {
	// Program data
	b, error := io.ReadAll(r)

	if error != nil {
		log.Println(error)
		return "", error
	}

	program := string(b)
	currentTime := time.Now().Format("2006-01-02 15:04:05")

	pd := ProgramData{
		Uid:         id,
		Program:     program,
		DateUpdated: currentTime,
		DType:       []string{"Program"},
	}

	uid, error := database.MakeMutationAdd(pd)

	if error != nil {
		log.Println(error)
		return "", error
	}

	return uid, nil
}

// AllPrograms returns all programs from dgraph
func AllPrograms(offset string, first string) ([]byte, error) {

	// Query
	q := `
	{
		programs(func: type(Program),offset: ` + offset + `, first: ` + first + `) {
			uid
			name
			date_created
			date_updated
			program
		  }
	}
	`
	//
	res, error := database.MakeQuery(q)

	if error != nil {
		log.Println(error)
		return nil, error
	}

	return res.Json, nil
}

// GetProgram return program by id from dgraph
func GetProgram(id string) ([]byte, error) {
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
	res, error := database.MakeQuery(q)

	if error != nil {
		log.Println(error)
		return nil, error
	}

	return res.Json, nil
}

// CountPrograms return number of programs from dgraph
func CountPrograms() ([]byte, error) {
	// Query
	q := `
	{
		programs_counts(func: type(Program)){
			count(uid)
  		}
	}
	`
	//
	res, error := database.MakeQuery(q)

	if error != nil {
		log.Println(error)
		return nil, error
	}

	return res.Json, nil
}

// DeleteProgram delete program by id from dgraph
func DeleteProgram(id string) error {

	error := database.MakeMutationDelete(id)

	if error != nil {
		log.Println(error)
		return error
	}

	return nil
}
