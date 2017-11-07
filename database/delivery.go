package database

type Delivery struct {
	*Model
	MessageID  string `db:"message_id"`
	Status     string `db:"status"`
	Latency    int    `db:"latency"`
	StatusCode string `db:"status_code"`
}

func NewDelivery(messageID string) (*Delivery, error) {
	return &Delivery{
		Model:     NewModel("del"),
		MessageID: messageID,
		Status:    "processing",
	}, nil
}
