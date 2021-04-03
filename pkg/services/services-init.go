package services

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type Services struct {
	dbClient *mongo.Client
}
