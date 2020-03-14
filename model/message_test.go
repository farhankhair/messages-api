package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewMessage(t *testing.T) {
	t.Run("Ensure ID and Body attached when create message", func(t *testing.T) {
		body := "Hello, World!"

		message := *NewMessage(body)

		assert.NotNil(t, message.GetID())
		assert.EqualValues(t, body, message.Body)
	})
}
