package database

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
)

type Message struct {
	*Model
	WebhookID string `db:"webhook_id"`
	Signature string `db:"signature"`
	Headers   []byte `db:"headers"`
	Payload   []byte `db:"payload"`
}

func NewMessage(webhookID string, payload []byte) (*Message, error) {
	signature := CalculateSignature(payload, []byte("s3cr3t"))
	headers, err := NewMessageHeaders(signature)

	if err != nil {
		return nil, err
	}

	return &Message{
		Model:     NewModel("msg"),
		Signature: signature,
		Headers:   headers,
		Payload:   payload,
		WebhookID: webhookID,
	}, nil
}

func NewMessageHeaders(signature string) ([]byte, error) {
	headers := struct {
		ContentType   string `json:"Content-Type"`
		UserAgent     string `json:"User-Agent"`
		XHubSignature string `json:"X-Hub-Signature"`
	}{
		ContentType:   "application/json",
		UserAgent:     "Webhulk/1.0",
		XHubSignature: "sha1=" + signature,
	}

	encoded, err := json.Marshal(headers)

	if err != nil {
		return nil, err
	}

	return encoded, nil
}

func CalculateSignature(payload, key []byte) string {
	h := hmac.New(sha1.New, key)
	h.Write(payload)

	return hex.EncodeToString(h.Sum(nil))
}
