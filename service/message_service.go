//go:generate mockgen -destination=./mock_service/mock_message_service.go -source=message_service.go

package service

import (
	"github.com/farhanramadhan/messages-api/constant"
	"github.com/farhanramadhan/messages-api/model"
	"github.com/farhanramadhan/messages-api/repository"

	mqtts "github.com/eclipse/paho.mqtt.golang"
)

type MessageService interface {
	GetAllMessages() []model.Message
	InsertMessage(body string) error
}

// MessageService :nodoc:
type messageService struct {
	repo      repository.MessageRepository
	publisher mqtts.Client
}

// NewMessageService is a constructor
func NewMessageService(repo repository.MessageRepository, publisher mqtts.Client) *messageService {
	service := &messageService{
		repo:      repo,
		publisher: publisher,
	}

	return service
}

// GetAllMessages :nodoc:
func (ms *messageService) GetAllMessages() []model.Message {
	result := ms.repo.GetAllMessages()

	return result
}

// InsertMessage :nodoc:
func (ms *messageService) InsertMessage(body string) error {
	message := *model.NewMessage(body)

	err := ms.repo.InsertMessage(message)

	if err != nil {
		return err
	}

	go func() {
		ms.publisher.Publish(constant.InsertMessageTopicMQTT, 0, false, body)
	}()

	return nil
}
