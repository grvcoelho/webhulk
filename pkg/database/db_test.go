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

		webhookStore, err := NewWebhookStore(db)
		assert.NoError(t, err)

		w1, err := NewWebhook("events", "https://receiver.com", true)
		assert.NoError(t, err)

		_, err = webhookStore.Store(w1)
		assert.NoError(t, err, "Should store a webhook")

		w2, ok := webhookStore.Retrieve(w1.ID)
		assert.True(t, ok)
		assert.Equal(t, w1.ID, w2.ID)
	})

	t.Run("MessageStore", func(t *testing.T) {
		db, _ := NewDatabase(&cfg.Database{
			Address: "postgres://postgres:postgres@database/webhulk?sslmode=disable",
		})

		webhookStore, _ := NewWebhookStore(db)
		messageStore, err := NewMessageStore(db)
		assert.NoError(t, err)

		w1, _ := NewWebhook("events", "https://receiver.com", true)
		w1, _ = webhookStore.Store(w1)

		m1, err := NewMessage(w1.ID, []byte(`{"event": "status_changed"}`))
		assert.NoError(t, err)

		_, err = messageStore.Store(m1)
		assert.NoError(t, err, "Should store a message")

		m2, ok := messageStore.Retrieve(m1.ID)
		assert.True(t, ok)
		assert.Equal(t, m1.ID, m2.ID)
	})
}
