package apiserver

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAPIServer_Health(t *testing.T) {
	config := NewConfig()
	s := New(config)

	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/health", nil)
	s.handleHealth().ServeHTTP(rec, req)
	assert.Equal(t, rec.Body.String(), "{\"status\":\"OK\"}\n")
}