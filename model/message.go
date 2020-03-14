package model

import (
	"github.com/google/uuid"
)

// Message struct
type Message struct {
	ID   uuid.UUID `json:"id"`
	Body string    `json:"body"`
}

// NewMessage is a constructor for Message
func NewMessage(body string) *Message {
	message := &Message{
		ID:   uuid.New(),
		Body: body,
	}
	return message
}

// GetID :nodoc:
func (m *Message) GetID() uuid.UUID {
	return m.ID
}

// GetBody :nodoc:
func (m *Message) GetBody() string {
	return m.Body
}
