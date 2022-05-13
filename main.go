package main

import (
	"net/http"

	"github.com/nelsinho10/learn-to-program-api/routes"
)

func main() {
	// Inicializando las rutas y el servidor
	r := routes.Router()
	http.ListenAndServe(":3000", r)

}
