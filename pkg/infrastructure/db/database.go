package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ctx = context.TODO()
var client *mongo.Client

func InitDb (dbUrl, databaseName string) (*mongo.Database, error) {
	clientOptions := options.Client().ApplyURI(dbUrl)
	if client == nil {
		client, err := mongo.Connect(ctx, clientOptions)
		if err != nil {
			return nil, err
		}

		err = client.Ping(ctx, nil)
		if err != nil {
			return nil, err
		}
		log.Println("Database connection established")
		return client.Database(databaseName), nil
	}
	return client.Database(databaseName), nil
}
