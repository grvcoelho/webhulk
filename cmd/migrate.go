package cmd

import (
	"github.com/mattes/migrate"
	_ "github.com/mattes/migrate/database/postgres"
	_ "github.com/mattes/migrate/source/file"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	cfg "github.com/grvcoelho/webhulk/pkg/config"
)

func init() {
	Webhulk.AddCommand(Migrate)
	Migrate.Flags().StringP("source", "s", "", "Source location of the migrations")
	Migrate.Flags().StringP("config", "c", "", "The webhulk configuration file path")
}

var Migrate = &cobra.Command{
	Use:   "migrate",
	Short: "Migrate database",
	Run:   ParseMigrate,
}

func ParseMigrate(cmd *cobra.Command, args []string) {
	configFile := getConfigFile(cmd)
	config, err := cfg.Load(configFile)

	if err != nil {
		log.WithFields(log.Fields{
			"path":  configFile,
			"error": err,
		}).Fatal("Failed reading configuration file")
		return
	}

	sourceFlag, _ := getFlag(cmd, "source")
	source := defaultTo(config.Database.Migrations, sourceFlag)
	config.Database.Migrations = source

	RunMigrate(config.Database)
}

func RunMigrate(conf *cfg.Database) {
	m, err := migrate.New(conf.Migrations, conf.Address)

	if err != nil {
		log.WithFields(log.Fields{
			"address":    conf.Address,
			"migrations": conf.Migrations,
			"error":      err,
		}).Fatal("Failed starting migrations")
		return
	}

	if err := m.Up(); err != nil {
		log.WithFields(log.Fields{
			"address":    conf.Address,
			"migrations": conf.Migrations,
			"error":      err,
		}).Fatal("Failed running migrations")
		return
	}
}
