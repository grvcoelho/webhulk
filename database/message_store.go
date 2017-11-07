package database

import (
	"sync"

	"github.com/jmoiron/sqlx"
)

type MessageStore struct {
	db    *sqlx.DB
	mutex sync.RWMutex
}

func NewMessageStore(db *sqlx.DB) (*MessageStore, error) {
	return &MessageStore{
		db:    db,
		mutex: sync.RWMutex{},
	}, nil
}

func (s *MessageStore) Store(m *Message) (*Message, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	query := `
		INSERT INTO messages
			(id, headers, payload, signature, webhook_id, created_at, updated_at)
		VALUES
			(:id, :headers, :payload, :signature, :webhook_id, :created_at, :updated_at)
	`

	_, err := s.db.NamedExec(query, m)

	if err != nil {
		return nil, err
	}

	return m, nil
}

func (s *MessageStore) Retrieve(id string) (*Message, bool) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	m := &Message{}

	query := `
		SELECT
			id, headers, payload, signature, webhook_id, created_at, updated_at
		FROM messages
		WHERE
			id = $1
	`

	err := s.db.Get(m, query, id)

	if err != nil {
		return nil, false
	}

	return m, true
}
