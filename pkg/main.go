package main

import (
	"fmt"
	"github.com/gate3/sport-odds/pkg/config"
	"log"
)

func main () {
	_, err := config.LoadEnvironmentVariables(".")
	if err != nil {
		log.Fatal("Error loading config:", err)
	}

	//Recover From Panic
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
}
