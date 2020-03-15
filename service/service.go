package service

import (
	"github.com/farhanramadhan/messages-api/model"
	"github.com/farhanramadhan/messages-api/mqtt"
	"github.com/farhanramadhan/messages-api/repository"
	"github.com/farhanramadhan/messages-api/repository/localdb"

	mqtts "github.com/eclipse/paho.mqtt.golang"
)

// MessageService :nodoc:
type MessageService struct {
	repo      repository.MessageRepository
	publisher mqtts.Client
}

// NewMessageService is a constructor
func NewMessageService() *MessageService {
	service := &MessageService{
		repo:      localdb.NewLocalDBRepo(),
		publisher: mqtt.Publisher(),
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

	go func() {
		ms.publisher.Publish("message-api-realtime", 0, false, body)
	}()

	return nil
}
