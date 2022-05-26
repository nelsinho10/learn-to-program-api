package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/nelsinho10/learn-to-program-api/routes"
)

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	addr := fmt.Sprintf(":%s", os.Getenv("PORT"))

	// Create a new router
	r := routes.Router()

	fmt.Printf("Starting server on %v\n", addr)
	http.ListenAndServe(addr, r)
}
