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

	// Recover From Panic
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
	log.Println("Sports data saved successfully!")

	odds := new([]bookmaker.SportOddsApiModel)
	err = bk.FetchFixtures("upcoming","uk", "h2h", odds)
	if err != nil {
		log.Fatalln("An error occured! Could not fetch odds " + err.Error())
	}

	_, err = db.SaveFixtures(odds, dbCon)
	if err != nil {
		log.Fatalln("An error occurred saving odds "+ err.Error())
	}
	log.Println("Fixtures data saved successfully!")
}
