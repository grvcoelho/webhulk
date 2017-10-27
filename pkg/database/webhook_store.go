package database

import (
	"sync"

	"github.com/jmoiron/sqlx"
)

type WebhookStore struct {
	db    *sqlx.DB
	mutex sync.RWMutex
}

func NewWebhookStore(db *sqlx.DB) (*WebhookStore, error) {
	return &WebhookStore{
		db:    db,
		mutex: sync.RWMutex{},
	}, nil
}

func (s *WebhookStore) Store(w *Webhook) (*Webhook, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	query := `
		INSERT INTO webhooks
			(id, name, url, enabled, created_at)
		VALUES
		(:id, :name, :url, :enabled, :created_at)
	`

	_, err := s.db.NamedExec(query, w)

	if err != nil {
		return nil, err
	}

	return w, nil
}

func (s *WebhookStore) Retrieve(id string) (*Webhook, bool) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	w := &Webhook{}

	query := `
		SELECT
			id, name, url, enabled, created_at
		FROM webhooks
		WHERE
			id = $1
	`

	err := s.db.Get(w, query, id)

	if err != nil {
		return nil, false
	}

	return w, true
}
