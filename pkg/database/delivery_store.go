package database

import (
	"sync"

	"github.com/jmoiron/sqlx"
)

type DeliveryStore struct {
	db    *sqlx.DB
	mutex sync.RWMutex
}

func NewDeliveryStore(db *sqlx.DB) (*DeliveryStore, error) {
	return &DeliveryStore{
		db:    db,
		mutex: sync.RWMutex{},
	}, nil
}

func (s *DeliveryStore) Store(d *Delivery) (*Delivery, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	query := `
		INSERT INTO deliveries
			(id, status, latency, status_code, message_id, created_at, updated_at)
		VALUES
			(:id, :status, :latency, :status_code, :message_id, :created_at, :updated_at)
	`

	_, err := s.db.NamedExec(query, d)

	if err != nil {
		return nil, err
	}

	return d, nil
}

func (s *DeliveryStore) Retrieve(id string) (*Delivery, bool) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	d := &Delivery{}

	query := `
		SELECT
			id, status, latency, status_code, message_id, created_at, updated_at
		FROM deliveries
		WHERE
			id = $1
	`

	err := s.db.Get(d, query, id)

	if err != nil {
		return nil, false
	}

	return d, true
}
