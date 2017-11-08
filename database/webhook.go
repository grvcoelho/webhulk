package database

type Webhook struct {
	*Model
	Name    string `db:"name"`
	URL     string `db:"url"`
	Enabled bool   `db:"enabled"`
}

func NewWebhook(name, url string, enabled bool) (*Webhook, error) {
	return &Webhook{
		Model:   NewModel("web"),
		Name:    name,
		URL:     url,
		Enabled: enabled,
	}, nil
}
