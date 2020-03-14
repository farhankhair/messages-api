package main

import (
	"fmt"

	"github.com/farhanramadhan/messages-api/service"
)

func main() {
	message := "Hello, World!"
	service := service.NewMessageService()
	service.InsertMessage(message)
	data := service.GetAllMessages()
	fmt.Println(data)
}
