package sqlstore

import (
	"database/sql"
	"github.com/dev-tim/message-board-api/internal/app/store"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

type Store struct {
	db                 *sql.DB
	messagesRepository store.IMessagesRepository
}

func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) Messages() store.IMessagesRepository {
	if s.messagesRepository == nil {
		s.messagesRepository = NewMessageRepository(s)
	}

	return s.messagesRepository
}
