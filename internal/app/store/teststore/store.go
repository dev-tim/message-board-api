package teststore

import (
	"github.com/dev-tim/message-board-api/internal/app/store"
)

type Store struct {
	messagesRepository store.IMessagesRepository
}

func New() *Store {
	return &Store{}
}

func (s *Store) Messages() store.IMessagesRepository {
	return NewMessageRepository(s)
}
