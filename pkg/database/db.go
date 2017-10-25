package database

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"

	cfg "github.com/grvcoelho/webhulk/pkg/config"
)

func NewDatabase(conf *cfg.Database) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", conf.Address)

	if err != nil {
		log.WithFields(log.Fields{
			"address": conf.Address,
			"error":   err,
		})

		return nil, err
	}

	return db, err
}
