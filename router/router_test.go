package router

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/farhanramadhan/messages-api/service/mock_service"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestAPI(t *testing.T) {
	t.Run("Get all message from API", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		req, err := http.NewRequest("GET", "/message", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()

		router := Router()

		mockService := mock_service.NewMockMessageService(ctrl)

		// mock call
		mockService.EXPECT().GetAllMessages()

		handler := http.HandlerFunc(NewAPI(mockService, router).getAllMessages)

		handler.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
		assert.Equal(t, "{\"data\":{\"length\":0,\"messages\":null}}", rr.Body.String())
	})

	t.Run("Insert a message", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		req, err := http.NewRequest("GET", "/message", nil)
		if err != nil {
			t.Fatal(err)
		}

		q := req.URL.Query()
		q.Add("message", "insertTest")

		router := Router()

		rr := httptest.NewRecorder()
		mockService := mock_service.NewMockMessageService(ctrl)

		// mock call
		mockService.EXPECT().InsertMessage(gomock.Any())

		handler := http.HandlerFunc(NewAPI(mockService, router).insertMessage)

		handler.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
		assert.Equal(t, "{\"data\":{\"message\":\"\",\"status\":\"Success\"}}", rr.Body.String())
	})
}
