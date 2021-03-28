package main

import (
	"fmt"
	"github.com/gate3/sport-odds/pkg/config"
	"github.com/gate3/sport-odds/pkg/infrastructure"
	"log"
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

	_, err = infrastructure.InitDb(cfg.MongodbUri)
}
