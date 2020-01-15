package db

import (
	"context"
	"fmt"

	"api/model"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type EntryRepository interface {
	GetAll() ([]model.Visitor, error)
	Create(model.Visitor) error
	SignOut(id primitive.ObjectID) (model.Visitor, error)
}

type EntryMongoDBRepository struct {
	ctx     context.Context
	mongodb *mongo.Database
}

func NewEntryRepository(ctx context.Context, mongodb *mongo.Database) EntryMongoDBRepository {
	return EntryMongoDBRepository{
		ctx:     ctx,
		mongodb: mongodb,
	}
}

func (repo EntryMongoDBRepository) Collection() *mongo.Collection {
	return repo.mongodb.Collection("entry")
}

func (repo EntryMongoDBRepository) GetAll() (response []model.Visitor, err error) {
	cursor, err := repo.Collection().Find(repo.ctx, bson.D{})

	if nil != err {
		return nil, fmt.Errorf("Error retrieving all: %v", err)
	}

	defer cursor.Close(repo.ctx)
	for cursor.Next(repo.ctx) {
		var result model.Visitor
		err := cursor.Decode(&result)
		if err != nil {
			return nil, fmt.Errorf("Error decoding cursor: %v", err)
		}
		response = append(response, result)
	}

	return
}

func (repo EntryMongoDBRepository) Create(visitor model.Visitor) (err error) {
	visitor.CreatedAt = time.Now()
	bytes, err := bson.Marshal(visitor)
	if err != nil {
		fmt.Println("Error marshaling", err)
	}

	res, err := repo.Collection().InsertOne(repo.ctx, bytes)

	fmt.Println(res)

	return
}

func (repo EntryMongoDBRepository) SignOut(id primitive.ObjectID) (visitor model.Visitor, err error) {
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
	}
	result := repo.Collection().FindOneAndUpdate(
		repo.ctx,
		bson.M{"_id": id},
		bson.M{
			"$set": bson.M{
				"sign_out_time": time.Now(),
			},
		},
		&opt,
	)

	if nil != result.Err() {
		return visitor, result.Err()
	}

	err = result.Decode(&visitor)

	return
}
