package repository

import (
	"github.com/farhanramadhan/messages-api/model"
)

// MessageRepository :nodoc:
type MessageRepository interface {
	GetAllMessages() []model.Message
	InsertMessage(message model.Message) error
}
