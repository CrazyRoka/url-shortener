package main

import (
	"flag"
	"fmt"

	log "github.com/sirupsen/logrus"
)

func main() {
	log.Info("Starting application. Reading flags")
	port := *flag.Uint("p", 8080, "port for application")
	log.WithFields(log.Fields{
		"port": port,
	}).Info("Finished flags reading")

	InitDatabase()
	r := CreateRoutes()

	log.WithField("port", port).Info("Running application")
	r.Run(fmt.Sprintf(":%d", port))
}
