package db

import (
	"github.com/gate3/sport-odds/pkg/common/bookmaker"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type Repository struct {
	db	*mongo.Database
}

type IRepository interface {
	SaveFixtures (*[]bookmaker.SportOddsApiModel) ([]interface{}, error)
	SaveSports (*[]bookmaker.SportApiModel) ([]interface{}, error)
}

func NewRepository (databaseUri, dbName string) *Repository {
	dbCon, err := InitDb(databaseUri, dbName)
	if err != nil {
		log.Fatalln("Could not establish a connection to the database")
	}

	return &Repository{
		db: dbCon,
	}
}
