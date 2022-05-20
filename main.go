package main

import (
	"net/http"

	"github.com/nelsinho10/learn-to-program-api/routes"
)

func main() {
	// Create a new router
	r := routes.Router()
	http.ListenAndServe(":3000", r)
}
