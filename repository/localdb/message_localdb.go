package localdb

import "github.com/farhanramadhan/messages-api/model"

// Repo is a temporary database
type repo struct {
	data []model.Message
}

// NewLocalDBRepo is a constructor
func NewLocalDBRepo() *repo {
	repo := &repo{
		data: make([]model.Message, 0),
	}
	return repo
}

// GetAllMessages is to get all messages
func (lc *repo) GetAllMessages() []model.Message {
	return lc.data
}

// InsertMessage is to insert a message
func (lc *repo) InsertMessage(message model.Message) error {
	lc.data = append(lc.data, message)
	return nil
}
