package db

import (
	"context"
	"fmt"
	"time"

	"api/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type EntryRepository interface {
	GetAll() ([]model.Visitor, error)
	Create(model.Visitor) error
}

type EntryMongoDBRepository struct {
	mongodb *mongo.Database
}

func NewEntryRepository(mongodb *mongo.Database) EntryMongoDBRepository {
	return EntryMongoDBRepository{
		mongodb: mongodb,
	}
}

func (repo EntryMongoDBRepository) Collection() *mongo.Collection {
	return repo.mongodb.Collection("entry")
}

func (repo EntryMongoDBRepository) GetAll() ([]model.Visitor, error) {

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	cursor, err := repo.Collection().Find(ctx, bson.D{})

	if nil != err {
		fmt.Println("Error retrieving all entries", err)
	}

	var response []model.Visitor
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		// var result bson.M
		var result model.Visitor
		err := cursor.Decode(&result)
		if err != nil {
			return nil, fmt.Errorf("Error decoding cursor", err)
		}
		response = append(response, result)
	}

	// response := []model.Visitor{}
	// for thing := range result {
	// 	visitor :=
	// 	response = append(response, model.Visitor{
	// 		id: thing.ID,
	// 	})

	// }

	return response, nil
}

func (repo EntryMongoDBRepository) Create(visitor model.Visitor) (err error) {

	bytes, err := bson.Marshal(visitor)
	if err != nil {
		fmt.Println("Error marshaling", err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	res, err := repo.Collection().InsertOne(ctx, bytes)

	fmt.Println(res)

	return err
}
