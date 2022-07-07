package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoRepository struct {
	db *mongo.Database
}

func NewMongoRepository(uri string) (*MongoRepository, error) {
	client, connErr := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
	if connErr != nil {
		return nil, connErr
	}
	return &MongoRepository{db: client.Database("platzi")}, nil
}
