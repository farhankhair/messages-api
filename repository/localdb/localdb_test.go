package localdb

import (
	"testing"

	"github.com/farhanramadhan/messages-api/model"
	"github.com/stretchr/testify/assert"
)

func TestRepo(t *testing.T) {
	repo := NewLocalDBRepo()

	body := "Hello, World!"

	t.Run("Ensure DB Empty when created", func(t *testing.T) {
		assert.Equal(t, 0, len(repo.data))
	})

	t.Run("Ensure to insert new message", func(t *testing.T) {
		message := *model.NewMessage(body)

		err := repo.InsertMessage(message)

		assert.Nil(t, err)
		assert.EqualValues(t, body, message.GetBody())
	})

	t.Run("Ensure to get all messages in database", func(t *testing.T) {
		result := repo.GetAllMessages()

		assert.EqualValues(t, 1, len(result))
		assert.EqualValues(t, body, result[0].GetBody())
	})
}
