package sqlstore_test

import (
	"github.com/dev-tim/message-board-api/internal/app/model"
	"github.com/dev-tim/message-board-api/internal/app/store/sqlstore"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var now = time.Now().UTC()
var testMessages = []model.Message{
	{
		Id:           "a1",
		Name:         "Testname",
		Email:        "foo1@bar.baz",
		Text:         "Some short tex",
		CreationTime: &now,
		CreatedAt:    nil,
		UpdatedAt:    nil,
	},
	{
		Id:           "a2",
		Name:         "Testname2",
		Email:        "foo2@bar.baz",
		Text:         "Some short tex",
		CreationTime: &now,
		CreatedAt:    nil,
		UpdatedAt:    nil,
	},
}

func TestMessagesRepository_Create(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseUrl)
	logger, _ := test.NewNullLogger()
	s := sqlstore.New(db, logger)
	defer teardown("messages")

	now := time.Now()
	createdMessage, err := s.Messages().Create(&model.Message{
		Id:           "a-b-d",
		Name:         "Testname",
		Email:        "foo@bar.baz",
		Text:         "Some short tex",
		CreationTime: &now,
		CreatedAt:    nil,
		UpdatedAt:    nil,
	})

	assert.NoError(t, err)
	assert.NotNil(t, createdMessage)
}

func TestMessagesRepository_Find_When_2_EntriesExist(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseUrl)
	logger, _ := test.NewNullLogger()
	s := sqlstore.New(db, logger)
	defer teardown("messages")

	for _, m := range testMessages {
		_, _ = s.Messages().Create(&m)
	}

	messages, err := s.Messages().FindLatest(5, 0)

	assert.NoError(t, err)
	assert.Equal(t, len(messages), 2)
	for idx, s := range messages {
		assert.Equal(t, s.Id, testMessages[idx].Id)
		assert.Equal(t, s.Name, testMessages[idx].Name)
		assert.Equal(t, s.Text, testMessages[idx].Text)
		assert.Equal(t, s.Email, testMessages[idx].Email)
		assert.Equal(t, s.CreationTime.String(), now.String())
		assert.NotEmpty(t, *s.CreatedAt)
		assert.NotEmpty(t, *s.UpdatedAt)
	}
}

func TestMessagesRepository_Find_When_NoEntries_Exist(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseUrl)
	logger, _ := test.NewNullLogger()
	s := sqlstore.New(db, logger)
	defer teardown("messages")

	messages, err := s.Messages().FindLatest(5, 0)

	assert.NoError(t, err)
	assert.Equal(t, len(messages), 0)
}

func TestMessagesRepository_Update(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseUrl)
	logger, _ := test.NewNullLogger()
	s := sqlstore.New(db, logger)
	defer teardown("messages")

	for _, m := range testMessages {
		_, _ = s.Messages().Create(&m)
	}

	update, err := s.Messages().Update("a2", "Updated text")
	if err != nil {
		t.Fatal(err)
	}
	t.Log("Updated entries ", update)

	messages, err := s.Messages().FindLatest(5, 0)

	assert.NoError(t, err)
	assert.Equal(t, len(messages), 2)

	for _, s := range messages {
		if s.Id == "a2" {
			assert.Equal(t, s.Text, "Updated text")
		} else {
			assert.NotEqual(t, s.Text, "Updated text")
		}
	}
}
