package main

import (
	"fmt"
	"math/rand"
	"time"

	log "github.com/sirupsen/logrus"
)

func main() {
	log.Info("Starting application")

	log.Info("Seeding random function")
	rand.Seed(time.Now().Unix())

	config, err := LoadConfig()
	if err != nil {
		log.WithError(err).Fatal("Error reading config")
	}

	log.WithFields(log.Fields{
		"config": config,
	}).Info("Finished loading environment variables")

	InitDatabase(config.DBHost)
	r := CreateRoutes()

	log.WithField("port", config.Port).Info("Running application")
	r.Run(fmt.Sprintf(":%d", config.Port))
}
