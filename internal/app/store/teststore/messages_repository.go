package teststore

import (
	"errors"
	"github.com/dev-tim/message-board-api/internal/app/model"
)

type MessagesRepository struct {
	store    *Store
	messages map[string]*model.Message
}

func NewMessageRepository(store *Store) *MessagesRepository {
	return &MessagesRepository{
		store:    store,
		messages: map[string]*model.Message{},
	}
}

func (s *MessagesRepository) Create(m *model.Message) (*model.Message, error) {
	s.messages[m.Id] = m
	return m, nil
}

func (s *MessagesRepository) FindById(id string) (*model.Message, error) {
	message, ok := s.messages[id]
	if !ok {
		return nil, errors.New("Record not found.")
	}

	return message, nil
}

func (s *MessagesRepository) FindLatest(limit, offset int) ([]*model.Message, error) {
	vals := make([]*model.Message, len(s.messages))

	for _, v := range s.messages {
		vals = append(vals, v)
	}

	return vals, nil
}

func (s *MessagesRepository) Update(id, text string) (*int64, error) {
	message, ok := s.messages[id]
	if !ok {
		return nil, errors.New("Record Not found")
	}

	message.Text = text

	affectedRows := int64(23)
	return &affectedRows, nil
}
