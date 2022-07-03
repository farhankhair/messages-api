package main

import (
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/farhanramadhan/messages-api/mqtt"
	"github.com/farhanramadhan/messages-api/repository/localdb"
	"github.com/farhanramadhan/messages-api/router"
	"github.com/farhanramadhan/messages-api/service"
)

func main() {
	startServer()
}

func startServer() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// initialize dependency
	publisher := mqtt.Publisher()

	repo := localdb.NewLocalDBRepo()

	service := service.NewMessageService(repo, publisher)

	muxRouter := router.Router()

	api := router.NewAPI(service, muxRouter)

	server := &http.Server{
		Handler:      api.Router,
		Addr:         ":" + port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	// start mqtt subscriber go routine
	go startMQTTSubscriber()

	log.Println("Starting server on port ", port)

	// start server
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
