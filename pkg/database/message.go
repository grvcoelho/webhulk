package database

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

type Message struct {
	*Model
	WebhookID string `db:"webhook_id"`
	Signature string `db:"signature"`
	Headers   []byte `db:"headers"`
	Payload   []byte `db:"payload"`
}

func NewMessage(webhookID string, payload []byte) (*Message, error) {
	m := &Message{
		Model:     NewModel("msg"),
		Headers:   []byte("{}"),
		Payload:   payload,
		WebhookID: webhookID,
	}

	m.Signature = m.CalculateSignature()

	return m, nil
}

func (m *Message) CalculateSignature() string {
	key := []byte("s3cr3t")
	h := hmac.New(sha256.New, key)
	h.Write(m.Payload)

	return hex.EncodeToString(h.Sum(nil))
}
