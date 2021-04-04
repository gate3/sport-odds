package db

import (
	"github.com/gate3/sport-odds/pkg/common/bookmaker"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {}

type IRepository interface {
	SaveFixtures (*[]bookmaker.SportOddsApiModel, *mongo.Database) (*mongo.InsertManyResult, error)
	SaveSports (*[]bookmaker.SportApiModel, *mongo.Database) (*mongo.InsertManyResult, error)
}

func NewRepository () *Repository {
	return &Repository{}
}
