package apiserver_test

import (
	"github.com/dev-tim/message-board-api/internal/app/apiserver"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestContextMiddleware(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/health", nil)

	handler := func() http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(200)
			assert.NotEmpty(t, r.Context().Value("requestId"))
		}
	}

	m := apiserver.ContextMiddleware(handler())
	m.ServeHTTP(rec, req)

}
