package database

import (
	"time"

	cuid "gopkg.in/lucsky/cuid.v1"
)

type Model struct {
	ID        string    `db:"id"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func NewModel(prefix string) *Model {
	now := time.Now()

	return &Model{
		ID:        newIDWithPrefix(prefix),
		CreatedAt: now,
		UpdatedAt: now,
	}
}

func newID() string {
	return cuid.New()
}

func newIDWithPrefix(prefix string) string {
	return prefix + "_" + newID()
}
