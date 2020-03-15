package router

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/farhanramadhan/messages-api/service"
)

// API :nodoc:
type API struct {
	Service *service.MessageService
	Router  *mux.Router
}

// NewAPI :nodoc:
func NewAPI() *API {
	api := API{
		Service: service.NewMessageService(),
		Router:  Router(),
	}

	// Route for sending message and get all messages
	api.Router.HandleFunc("/message/{message}", api.insertMessage).Methods("GET")
	api.Router.HandleFunc("/message", api.getAllMessages).Methods("GET")

	return &api
}

// Router :nodoc:
func Router() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/", welcome)

	return r
}

func (a *API) insertMessage(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	message := params["message"]

	err := a.Service.InsertMessage(message)
	if err != nil {
		handleError(w, NewErrorNoMessage(400))
		return
	}

	var data struct {
		Data struct {
			Message string `json:"message"`
			Status  string `json:"status"`
		} `json:"data"`
	}

	data.Data.Message = message
	data.Data.Status = "Success"

	handleJSONResponse(w, data)
}

func (a *API) getAllMessages(w http.ResponseWriter, r *http.Request) {
	messages := a.Service.GetAllMessages()

	var data struct {
		Data struct {
			Length   int `json:"length"`
			Messages []struct {
				MessageID string `json:"id_message"`
				Body      string `json:"body"`
			} `json:"messages"`
		} `json:"data"`
	}

	for _, v := range messages {
		var message struct {
			MessageID string `json:"id_message"`
			Body      string `json:"body"`
		}

		message.MessageID = v.GetID().String()
		message.Body = v.GetBody()

		data.Data.Messages = append(data.Data.Messages, message)
	}

	data.Data.Length = len(messages)

	handleJSONResponse(w, data)
}
