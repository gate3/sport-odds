package services

import (
	"github.com/gate3/sport-odds/pkg/common/bookmaker"
	"github.com/gate3/sport-odds/pkg/config"
	"github.com/gate3/sport-odds/pkg/infrastructure/db"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type Services struct {
	Db 			*mongo.Database
	Dao			*db.Dao
	EnvVars 	config.EnvVariables
	Bookmaker 	*bookmaker.BookmakersApi
}

func NewServices (cfg config.EnvVariables) *Services {

	dbCon, err := db.InitDb(cfg.DatabaseUri, cfg.DatabaseName)
	if err != nil {
		log.Fatalln("Could not establish a connection to the database")
	}

	bk, err := bookmaker.NewBookmakerApi(cfg.OddsApiBaseUrl, cfg.OddsApiKey)
	if err != nil {
		log.Fatalln(err.Error())
	}

	dao := db.NewDao()

	return &Services{
		Db: dbCon,
		Bookmaker: bk,
		EnvVars: cfg,
		Dao: dao,
	}
}
