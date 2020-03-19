package apiserver_test

import (
	"github.com/dev-tim/message-board-api/internal/app/apiserver"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestBasicAuth_When_Creds_AreWrong(t *testing.T) {

	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/health", nil)
	req.SetBasicAuth("username", "false")

	m := apiserver.BasicAuth(rec, req, "Pass password")
	assert.False(t, m)
	assert.Equal(t, rec.Result().StatusCode, 401)
}

func TestBasicAuth_When_Creds_AreCorrect(t *testing.T) {

	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/health", nil)
	req.SetBasicAuth("test", "test")

	m := apiserver.BasicAuth(rec, req, "Pass password")
	assert.True(t, m)
	assert.Equal(t, rec.Result().StatusCode, 200)
}
