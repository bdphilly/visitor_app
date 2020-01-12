package db

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConfigureDB(ctx context.Context) (*mongo.Database, error) {

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if nil != err {
		return nil, fmt.Errorf("Error connecting to mongo: %v", err)
	}

	err = client.Connect(ctx)
	if nil != err {
		return nil, fmt.Errorf("Mongo failure connecting to context: %v", err)
	}

	entryDB := client.Database("entry")
	return entryDB, nil
}
