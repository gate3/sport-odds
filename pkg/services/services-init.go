package services

import (
	"github.com/gate3/sport-odds/pkg/common/bookmaker"
	"github.com/gate3/sport-odds/pkg/config"
	"github.com/gate3/sport-odds/pkg/infrastructure/db"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

var services bookmaker.IBookmakerApi
var repo db.IRepository

func init () {
	services = bookmaker.BookmakersApi{}
	repo = db.Repository{}
}

type Services struct {
	Db 			*mongo.Database
	Repository	db.IRepository
	EnvVars 	config.EnvVariables
	Bookmaker 	bookmaker.IBookmakerApi
}

func NewServices (cfg config.EnvVariables) *Services {
	services, err := bookmaker.NewBookmakerApi(cfg.OddsApiBaseUrl, cfg.OddsApiKey)
	if err != nil {
		log.Fatalln(err.Error())
	}

	repo := db.NewRepository(cfg.DatabaseUri, cfg.DatabaseName)

	return &Services{
		Bookmaker: services,
		EnvVars: cfg,
		Repository: repo,
	}
}
