package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	cfg "github.com/grvcoelho/webhulk/pkg/config"
	"github.com/grvcoelho/webhulk/pkg/server"
)

func init() {
	Webhulk.AddCommand(Server)
	Server.Flags().StringP("config", "c", "", "The webhulk configuration file path")
}

var Server = &cobra.Command{
	Use:   "server",
	Short: "Start Webhulk server",
	Run:   ParseServer,
}

func ParseServer(cmd *cobra.Command, args []string) {
	configFile := getConfigFile(cmd)
	config, err := cfg.Load(configFile)

	if err != nil {
		log.WithFields(log.Fields{
			"path":  configFile,
			"error": err,
		}).Fatal("Failed reading configuration file")
		return
	}

	RunServer(config)
}

func RunServer(config *cfg.Configuration) {
	log.WithFields(log.Fields{
		"config": config.Server,
	}).Info("Starting Webhulk server")

	s, err := server.New(config)

	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Fatal("Failed starting server")
		return
	}

	s.Start()
}
