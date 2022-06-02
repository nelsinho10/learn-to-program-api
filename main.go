package main

import (
	"fmt"
	"net/http"

	"github.com/nelsinho10/learn-to-program-api/helpers"
	"github.com/nelsinho10/learn-to-program-api/routes"
)

func main() {

	// Port to listen on
	addr := fmt.Sprintf(":%s", helpers.GetEnv("PORT", "3000"))

	// Create a new router
	r := routes.Router()

	fmt.Printf("Starting server on %v\n", addr)
	http.ListenAndServe(addr, r)
}
