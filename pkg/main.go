package main

import (
	"fmt"
	"github.com/gate3/sport-odds/pkg/services"
	"log"

	"github.com/gate3/sport-odds/pkg/config"
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

	srv := services.NewServices(cfg)

	_, err = srv.SaveAllSportsRecords()
	if err != nil {
		log.Fatal("Error loading config:", err)
	}

	_, err = srv.SaveUpcomingFixtureRecords()
	if err != nil {
		log.Fatal("Error loading config:", err)
	}
}
