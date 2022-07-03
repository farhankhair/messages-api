//go:generate mockgen -destination=./mock_repository/mock_message_repository.go -source=message_repository.go
package repository

import (
	"github.com/farhanramadhan/messages-api/model"
)

// MessageRepository :nodoc:
type MessageRepository interface {
	GetAllMessages() []model.Message
	InsertMessage(message model.Message) error
}
