package database

import "time"

type Webhook struct {
	*Model
	Name    string `db:"name"`
	URL     string `db:"url"`
	Enabled bool   `db:"enabled"`
}

func NewWebhook(name, url string, enabled bool) (*Webhook, error) {
	return &Webhook{
		Model: &Model{
			ID:        newIDWithPrefix("web"),
			CreatedAt: time.Now(),
		},
		Name:    name,
		URL:     url,
		Enabled: enabled,
	}, nil
}
