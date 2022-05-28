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

// dgraphClient returns a dgraph client instance
func dgraphClient() *dgo.Dgraph {
	// Load .env file
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

// MakeMutation makes a mutation
func MakeMutation(data any) string {
	var uid string
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

	for _, v := range res.Uids {
		uid = v
	}

	return uid

}

// MakeQuery makes a query and returns the result
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
