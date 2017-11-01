package database

type Message struct {
	*Model
	WebhookID string `db:"webhook_id"`
	Signature string `db:"signature"`
	Headers   []byte `db:"headers"`
	Payload   []byte `db:"payload"`
}

func NewMessage(webhookID string, payload []byte) (*Message, error) {
	return &Message{
		Model:     NewModel("msg"),
		Headers:   []byte("{}"),
		Payload:   payload,
		WebhookID: webhookID,
	}, nil
}
