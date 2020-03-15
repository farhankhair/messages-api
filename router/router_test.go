package router

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAPI(t *testing.T) {
	t.Run("Get all message from API", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/message", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(NewAPI().getAllMessages)

		handler.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
		assert.Equal(t, "{\"data\":{\"length\":0,\"messages\":null}}", rr.Body.String())
	})

	t.Run("Insert a message", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/message", nil)
		if err != nil {
			t.Fatal(err)
		}

		q := req.URL.Query()
		q.Add("message", "insertTest")

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(NewAPI().insertMessage)

		handler.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
		assert.Equal(t, "{\"data\":{\"message\":\"\",\"status\":\"Success\"}}", rr.Body.String())
	})
}
