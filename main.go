package main

import (
	"os"

	"github.com/grvcoelho/webhulk/cmd"
	log "github.com/sirupsen/logrus"
)

func main() {
	if err := cmd.Webhulk.Execute(); err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error()

		os.Exit(1)
	}
}
