package localdb

import "github.com/farhanramadhan/messages-api/model"

// Repo is a temporary database
type Repo struct {
	data []model.Message
}

// NewLocalDBRepo is a constructor
func NewLocalDBRepo() *Repo {
	repo := &Repo{
		data: make([]model.Message, 0),
	}
	return repo
}

// GetAllMessages is to get all messages
func (lc *Repo) GetAllMessages() []model.Message {
	return lc.data
}

// InsertMessage is to insert a message
func (lc *Repo) InsertMessage(message model.Message) error {
	lc.data = append(lc.data, message)
	return nil
}
