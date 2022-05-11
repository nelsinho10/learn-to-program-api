package configs

/*
ctx := context.TODO()

	dgraphClient := configs.NewClient()
	txn := dgraphClient.NewTxn()
	defer txn.Commit(ctx)

	url := fmt.Sprintf("https://example.com/%v", time.Now().UnixNano())

	link := models.Link{
		URL:   url,
		DType: []string{"Link"},
	}

	lb, err := json.Marshal(link)
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

*/
