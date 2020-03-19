package apiserver_test

import (
	"bytes"
	"encoding/json"
	"github.com/dev-tim/message-board-api/internal/app/apiserver"
	"github.com/dev-tim/message-board-api/internal/app/model"
	"github.com/dev-tim/message-board-api/internal/app/store/teststore"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAPIServer_Health(t *testing.T) {
	logger, _ := test.NewNullLogger()
	server := apiserver.New(&apiserver.Config{BindAddress: ":8080"}, nil, logger)

	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/health", nil)
	server.HandleHealth().ServeHTTP(rec, req)
	assert.Equal(t, rec.Body.String(), "{\"status\":\"OK\"}\n")
}

func TestAPIServer_TestCreate_GetAll(t *testing.T) {
	logger, _ := test.NewNullLogger()
	testStore := teststore.New()
	server := apiserver.New(&apiserver.Config{BindAddress: ":8080"}, testStore, logger)

	request := apiserver.CreateMessageBodyV1ClientRequest{
		Id:    "a",
		Name:  "first",
		Email: "foo@baz",
		Text:  "Test1",
	}
	jsonBytes, _ := json.Marshal(request)

	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/public/v1/messages", bytes.NewReader(jsonBytes))
	req.Header.Set("Content-Type", "application/json")

	server.HandlePostMessage().ServeHTTP(rec, req)

	assert.Equal(t, rec.Result().StatusCode, 201)

	rec = httptest.NewRecorder()
	req, _ = http.NewRequest(http.MethodGet, "/private/v1/messages", nil)
	req.SetBasicAuth("test", "test")
	server.HandleGetMessages().ServeHTTP(rec, req)

	assert.Equal(t, rec.Result().StatusCode, 200)
	m := make([]*model.Message, 0)

	_ = json.Unmarshal(rec.Body.Bytes(), &m)
	//assert.Equal(t, rec.Body.String(), "")
	assert.Equal(t, m[0].Id, "a")
	assert.Equal(t, m[0].Name, "first")
}

func TestAPIServer_Create_Update_Flow(t *testing.T) {
	logger, _ := test.NewNullLogger()
	testStore := teststore.New()
	app := apiserver.New(&apiserver.Config{BindAddress: ":8080"}, testStore, logger)
	app.ConfigureRouter()

	server := httptest.NewServer(app.Router)

	request := apiserver.CreateMessageBodyV1ClientRequest{
		Id:    "a",
		Name:  "first",
		Email: "foo@baz",
		Text:  "Test1",
	}
	jsonBytes, _ := json.Marshal(request)

	req, _ := http.NewRequest(http.MethodPost, server.URL+"/private/v1/messages", bytes.NewReader(jsonBytes))
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth("test", "test")

	rec, _ := server.Client().Do(req)
	assert.Equal(t, rec.StatusCode, 201)

	patchRequestJson := apiserver.PatchMessageBodyV1ClientRequest{
		Text: "ChangedText",
	}
	jsonBytes, _ = json.Marshal(patchRequestJson)

	req, _ = http.NewRequest(http.MethodPatch, server.URL+"/private/v1/messages/a", bytes.NewReader(jsonBytes))
	req.SetBasicAuth("test", "test")

	rec, _ = server.Client().Do(req)

	assert.Equal(t, rec.StatusCode, 200)

	req, _ = http.NewRequest(http.MethodGet, server.URL+"/private/v1/messages/a", nil)
	req.SetBasicAuth("test", "test")

	rec, _ = server.Client().Do(req)

	assert.Equal(t, rec.StatusCode, 200)

	m := &model.Message{}
	assert.Equal(t, rec.StatusCode, 200)

	_ = json.NewDecoder(rec.Body).Decode(&m)
	assert.Equal(t, m.Id, "a")
	assert.Equal(t, m.Text, "ChangedText")
}
