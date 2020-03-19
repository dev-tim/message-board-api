package sqlstore

import (
	"database/sql"
	"github.com/dev-tim/message-board-api/internal/app/store"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

type Store struct {
	db                 *sql.DB
	logger             *logrus.Logger
	messagesRepository store.IMessagesRepository
}

func New(db *sql.DB, logger *logrus.Logger) *Store {
	return &Store{
		db:     db,
		logger: logger,
	}
}

func (s *Store) Messages() store.IMessagesRepository {
	if s.messagesRepository == nil {
		s.messagesRepository = NewMessageRepository(s)
	}

	return s.messagesRepository
}
