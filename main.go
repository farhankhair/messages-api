package main

import (
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/farhanramadhan/messages-api/mqtt"
	"github.com/farhanramadhan/messages-api/router"
)

func main() {
	startServer()
}

func startServer() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	api := router.NewAPI()

	server := &http.Server{
		Handler:      api.Router,
		Addr:         ":" + port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	go startMQTTSubscriber()

	log.Println("Starting server on port ", port)

	log.Fatal(server.ListenAndServe())
}

func startMQTTSubscriber() {
	log.Println("Starting MQTT Subscriber")
	cloudMQTT := os.Getenv("CLOUDMQTT_URL")
	if cloudMQTT == "" {
		cloudMQTT = mqtt.LocalAddressMQTT()
	}

	uri, err := url.Parse(cloudMQTT)
	if err != nil {
		log.Fatal(err)
	}

	topic := uri.Path[1:len(uri.Path)]
	if topic == "" {
		topic = "message-api-realtime"
	}

	go mqtt.Listen(uri, topic)
}
