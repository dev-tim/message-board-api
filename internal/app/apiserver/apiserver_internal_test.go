package apiserver

import (
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAPIServer_Health(t *testing.T) {
	logger, _ := test.NewNullLogger()
	server := New(&Config{BindAddress: ":8080"}, nil, logger)

	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/health", nil)
	server.HandleHealth().ServeHTTP(rec, req)
	assert.Equal(t, rec.Body.String(), "{\"status\":\"OK\"}\n")
}
