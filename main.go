package main

import (
	"fmt"

	"github.com/aarjan/kafka/cmd"
	"github.com/aarjan/kafka/config"
	log "github.com/sirupsen/logrus"
)

func main() {
	// Load config
	c := config.Load()

	// Set application level logs
	if c.Debug {
		log.SetLevel(log.DebugLevel)
	}
	if c.LogFormatJSON {
		log.SetFormatter(&log.JSONFormatter{})
	}

	fmt.Println(c)
	cmd.Execute(c)
}
