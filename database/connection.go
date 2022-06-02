package database

import (
	"context"
	"encoding/json"
	"log"

	"github.com/dgraph-io/dgo/v200"
	"github.com/dgraph-io/dgo/v200/protos/api"
	"github.com/nelsinho10/learn-to-program-api/helpers"
)

// dgraphClient returns a dgraph client instance
func dgraphClient() (*dgo.Dgraph, error) {

	endpoint := helpers.GetEnv("DB_ENDPOINT", "http://localhost:8080/")
	key := helpers.GetEnv("DB_KEY", "")

	conn, err := dgo.DialSlashEndpoint(endpoint, key)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return dgo.NewDgraphClient(api.NewDgraphClient(conn)), nil
}

// MakeMutationAdd makes a mutation to add a new node
func MakeMutationAdd(data any) (string, error) {
	var uid string
	ctx := context.Background()

	client, err := dgraphClient()

	if err != nil {
		log.Println(err)
		return "", err
	}

	txn := client.NewTxn()
	defer txn.Commit(ctx)

	lb, err := json.Marshal(data)
	if err != nil {
		log.Println("Error marshalling data: ", err)
		return "", err
	}

	mu := &api.Mutation{
		SetJson: lb,
	}

	res, err := txn.Mutate(ctx, mu)
	if err != nil {
		log.Println("Error mutating data: ", err)
		return "", err
	}

	for _, v := range res.Uids {
		uid = v
	}

	return uid, nil
}

// MakeMutationDelete makes a mutation to delete a node
func MakeMutationDelete(uid string) error {
	ctx := context.Background()

	client, err := dgraphClient()

	if err != nil {
		log.Println(err)
		return err
	}

	txn := client.NewTxn()
	defer txn.Commit(ctx)

	d := map[string]string{"uid": uid}
	pb, err := json.Marshal(d)

	if err != nil {
		log.Println("Error marshalling data: ", err)
		return err
	}

	mu := &api.Mutation{
		CommitNow:  true,
		DeleteJson: pb,
	}

	_, err = txn.Mutate(ctx, mu)
	if err != nil {
		log.Println("Error mutating data: ", err)
		return err
	}

	return nil
}

// MakeQuery makes a query and returns the result
func MakeQuery(query string) (*api.Response, error) {
	ctx := context.Background()

	client, err := dgraphClient()

	if err != nil {
		log.Println(err)
		return nil, err
	}

	txn := client.NewTxn()
	defer txn.Discard(ctx)

	res, err := txn.Query(ctx, query)
	if err != nil {
		log.Println("Error querying data: ", err)
		return nil, err
	}
	return res, nil
}
