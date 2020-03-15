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

// GetID is a getter for message ID
func (m *Message) GetID() uuid.UUID {
	return m.ID
}

// GetBody is a getter for message body
func (m *Message) GetBody() string {
	return m.Body
}
