package router

import (
	"encoding/json"
	"net/http"
)

// ErrorResponse :nodoc:
type ErrorResponse struct {
	Status  int
	Message string
}

func handleJSONResponse(w http.ResponseWriter, v interface{}) {
	message, err := json.Marshal(v)

	if err != nil {
		handleError(w, NewErrorNoMessage(500))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(message)
}

func handleError(w http.ResponseWriter, err ErrorResponse) {
	var body struct {
		Message string `json:"message"`
	}

	body.Message = err.Message

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(err.Status)
	json.NewEncoder(w).Encode(body)
}

// NewErrorNoMessage :nodoc:
func NewErrorNoMessage(status int) ErrorResponse {
	return ErrorResponse{
		Status: status,
	}
}

type stringResponse struct {
	Data string `json:"data"`
}

func welcome(w http.ResponseWriter, r *http.Request) {
	hello := "Hello, World!"

	data := stringResponse{
		Data: hello,
	}

	handleJSONResponse(w, data)
}
