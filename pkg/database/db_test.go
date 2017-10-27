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

	t.Run("WebhookStore", func(t *testing.T) {
		db, _ := NewDatabase(&cfg.Database{
			Address: "postgres://postgres:postgres@database/webhulk?sslmode=disable",
		})

		store, err := NewWebhookStore(db)
		assert.NoError(t, err)

		w1, err := NewWebhook("events", "https://receiver.com", true)
		assert.NoError(t, err)

		_, err = store.Store(w1)
		assert.NoError(t, err, "Should store a webhook")

		w2, ok := store.Retrieve(w1.ID)
		assert.True(t, ok)
		assert.Equal(t, w1.ID, w2.ID)
	})
}
