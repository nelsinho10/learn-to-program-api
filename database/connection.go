package database

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/dgraph-io/dgo/v200"
	"github.com/dgraph-io/dgo/v200/protos/api"
	"github.com/joho/godotenv"
)

// dgraphClient realizar la conexion a la base de datos
func dgraphClient() *dgo.Dgraph {
	// Inicializar godotenv para obtener variables de entorno
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Obtener variables de entorno
	endpoint := os.Getenv("DB_ENDPOINT")
	key := os.Getenv("DB_KEY")

	// Conectandose a la base de datos
	conn, err := dgo.DialSlashEndpoint(endpoint, key)
	if err != nil {
		log.Fatal(err)
	}

	// Retornando el cliente de Dgraph
	return dgo.NewDgraphClient(api.NewDgraphClient(conn))
}

// MakeMutation realizar las mutaciones a la base de datos
func MakeMutation(data any) {
	ctx := context.Background()

	txn := dgraphClient().NewTxn()
	defer txn.Commit(ctx)

	lb, err := json.Marshal(data)
	if err != nil {
		log.Fatal("failed to marshal ", err)
	}

	mu := &api.Mutation{
		SetJson: lb,
	}

	res, err := txn.Mutate(ctx, mu)
	if err != nil {
		log.Fatal("failed to mutate ", err)
	}

	print("res: %v", res)
}

// MakeQuery realizar consultas a la base de datos
func MakeQuery(query string) *api.Response {
	ctx := context.Background()

	txn := dgraphClient().NewTxn()
	defer txn.Discard(ctx)

	res, err := txn.Query(ctx, query)
	if err != nil {
		log.Fatal(err)
	}
	return res
}
