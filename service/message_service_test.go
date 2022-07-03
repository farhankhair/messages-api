//go:generate mockgen -destination=./mock_service/mock_message_service.go -source=message_service.go

package service

import (
	"reflect"
	"testing"

	mqtts "github.com/eclipse/paho.mqtt.golang"
	"github.com/farhanramadhan/messages-api/model"
	"github.com/farhanramadhan/messages-api/repository"
	"github.com/farhanramadhan/messages-api/repository/mock_repository"
	"github.com/golang/mock/gomock"
)

func Test_messageService_GetAllMessages(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repository.NewMockMessageRepository(ctrl)
	type fields struct {
		repo      repository.MessageRepository
		publisher mqtts.Client
	}

	tests := []struct {
		name   string
		fields fields
		mocks  func()
		want   []model.Message
	}{
		{
			name: "success",
			fields: fields{
				repo: mockRepo,
			},
			mocks: func() {
				mockRepo.EXPECT().GetAllMessages().Return([]model.Message{})
			},
			want: []model.Message{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mocks()

			ms := &messageService{
				repo:      tt.fields.repo,
				publisher: tt.fields.publisher,
			}

			if got := ms.GetAllMessages(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("messageService.GetAllMessages() = %v, want %v", got, tt.want)
			}
		})
	}
}
