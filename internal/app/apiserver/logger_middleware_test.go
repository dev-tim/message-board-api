package apiserver_test

import (
	"context"
	"github.com/dev-tim/message-board-api/internal/app/apiserver"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLoggingMiddleware(t *testing.T) {
	logger, hook := test.NewNullLogger()

	rec := httptest.NewRecorder()
	ctx := context.WithValue(context.Background(), "requestId", "test")
	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, "/health", nil)

	m := apiserver.LoggingMiddleware(logger)
	m(HandleTest()).ServeHTTP(rec, req)

	assert.Equal(t, hook.LastEntry().Message, "[requestId] test [Resource] GET /health - [statusCode] - 200 [latency] 0 - agent ")
}

func HandleTest() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	}
}
