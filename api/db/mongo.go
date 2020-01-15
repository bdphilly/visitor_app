package db

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConfigureDB(ctx context.Context) (*mongo.Database, error) {

	host := os.Getenv("DB_HOST")
	if host == "" {
		// when testing
		host = "127.0.0.1:27017"
	}

	client, err := mongo.NewClient(options.Client().ApplyURI(fmt.Sprintf("mongodb://%v", host)))
	if nil != err {
		return nil, fmt.Errorf("Error connecting to mongo: %v", err)
	}

	err = client.Connect(ctx)
	if nil != err {
		return nil, fmt.Errorf("Mongo failure connecting to context: %v", err)
	}

	entryDB := client.Database("visitApp")
	if host == "" {
		// when testing
		entryDB = client.Database("test_visitApp")
	}
	return entryDB, nil
}
