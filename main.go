package main

import (
	"log"
	"net/http"
	"os"
	"time"

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

	log.Println("Starting server on port ", port)

	log.Fatal(server.ListenAndServe())
}
