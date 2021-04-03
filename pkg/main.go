package main

import (
	"fmt"
	"log"

	"github.com/gate3/sport-odds/pkg/common/bookmaker"
	"github.com/gate3/sport-odds/pkg/config"
	"github.com/gate3/sport-odds/pkg/infrastructure/db"
)

func main () {
	cfg, err := config.LoadEnvironmentVariables(".")
	if err != nil {
		log.Fatal("Error loading config:", err)
	}

	//Recover From Panic
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()

	dbCon, err := db.InitDb(cfg.DatabaseUri, cfg.DatabaseName)
	if err != nil {
		log.Fatalln("Could not establish a connection to the database")
	}

	bk, err := bookmaker.NewBookmakerApi(cfg.OddsApiBaseUrl, cfg.OddsApiKey)
	if err != nil {
		log.Fatalln(err.Error())
	}

	sp := new([]bookmaker.SportApiModel)
	err = bk.FetchSports(sp)

	_, err = db.SaveSports(sp, dbCon)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Sports collection saved successfully!")
}
