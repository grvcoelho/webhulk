package database

import (
	"testing"

	cfg "github.com/grvcoelho/webhulk/pkg/config"
	"github.com/stretchr/testify/assert"
)

func TestDatabase(t *testing.T) {
	t.Run("NewDatabase", func(t *testing.T) {
		db, err := NewDatabase(&cfg.Database{
			Address: "postgres://postgres:postgres@database/webhulk?sslmode=disable",
		})

		assert.NoError(t, err, "Should create a database instance")

		err = db.Ping()

		assert.NoError(t, err, "Should connect to the database")
	})
}
