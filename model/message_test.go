package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewMessage(t *testing.T) {
	body := "Hello, World!"

	message := *NewMessage(body)

	assert.NotNil(t, message.GetID())
	assert.EqualValues(t, body, message.GetBody())
}
