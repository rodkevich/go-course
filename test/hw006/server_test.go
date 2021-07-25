package hw006

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/rodkevich/go-course/homework/hw006/task01"
	"github.com/stretchr/testify/assert"
)

func TestEchoServer(t *testing.T) {
	t.Run("any request returning JSON with headers",
		func(t *testing.T) {
			e := task01.NewEchoServer("localhost:9080")
			handler := http.HandlerFunc(e.ShowHeaders)
			assert.HTTPStatusCode(t, handler, "GET", "/anything/you?want", nil, 200)
			assert.HTTPStatusCode(t, handler, "POST", "/want?you?or?not", nil, 200)
			request, _ := http.NewRequest(
				http.MethodGet,
				"/anything/you?want",
				nil,
			)
			response := httptest.NewRecorder()
			e.ShowHeaders(response, request)
			anythingYouWant := response.Body.String()
			var reqHeaders = []string{"host", "user_agent", "request_uri", "headers"}
			for _, h := range reqHeaders {
				assert.Contains(t, anythingYouWant, h)
			}
		},
	)
}
