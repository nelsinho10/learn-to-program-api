package configs

import (
	"log"
	"os"

	"github.com/dgraph-io/dgo/v200"
	"github.com/dgraph-io/dgo/v200/protos/api"
	"github.com/joho/godotenv"
)

// NewClient Se encarga de realizar la conexion a la base de datos
func NewClient() *dgo.Dgraph {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	endpoint := os.Getenv("DB_ENDPOINT")
	key := os.Getenv("DB_KEY")

	conn, err := dgo.DialSlashEndpoint(endpoint, key)
	if err != nil {
		log.Fatal(err)
	}

	return dgo.NewDgraphClient(api.NewDgraphClient(conn))
}
