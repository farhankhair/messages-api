package service

import (
	"github.com/farhanramadhan/messages-api/model"
	"github.com/farhanramadhan/messages-api/repository"
	"github.com/farhanramadhan/messages-api/repository/localdb"
)

// MessageService :nodoc:
type MessageService struct {
	repo repository.MessageRepository
}

// NewMessageService is a constructor
func NewMessageService() *MessageService {
	service := &MessageService{
		repo: localdb.NewLocalDBRepo(),
	}

	return service
}

// GetAllMessages :nodoc:
func (ms *MessageService) GetAllMessages() []model.Message {
	result := ms.repo.GetAllMessages()

	return result
}

// InsertMessage :nodoc:
func (ms *MessageService) InsertMessage(body string) error {
	message := *model.NewMessage(body)

	err := ms.repo.InsertMessage(message)

	if err != nil {
		return err
	}

	return nil
}
