package store

import (
	"database/sql"
	"fmt"
	"github.com/dev-tim/message-board-api/internal/app"
	"github.com/dev-tim/message-board-api/internal/app/common"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"time"
)

type Store struct {
	config             *Config
	db                 *sql.DB
	messagesRepository *MessagesRepository
}

func New(storeConfig *Config) *Store {
	return &Store{
		config: storeConfig,
	}
}

func (s *Store) Messages() *MessagesRepository {
	if s.messagesRepository == nil {
		s.messagesRepository = NewMessageRepository(s)
	}

	return s.messagesRepository
}

func (s *Store) Open() error {
	logger := common.GetLogger()

	db, err := sql.Open("postgres", s.config.DbUrl)
	if err != nil {
		logger.Error("Failed to open connection to db", err)
		return err
	}

	// We want to limit connection pool or use pgbouncer. Here I decided to look at Go connection pooling mechanism
	db.SetMaxIdleConns(s.config.MaxIdleConnections)
	db.SetMaxOpenConns(s.config.MaxOpenConnections)

	// Set the maximum lifetime of a connection to 1 hour. Setting it to 0
	// means that there is no maximum lifetime and the connection is reused
	// forever (which is the default behavior).
	db.SetConnMaxLifetime(time.Hour)

	if err = db.Ping(); err != nil {
		logger.Error("Failed to ping db with opened connection", err)
		return err
	}

	logger.Info("DB Connection has been established, Success!")
	s.db = db

	return nil
}

func (s *Store) Migrate() error {
	logger := common.GetLogger()

	driver, err := postgres.WithInstance(s.db, &postgres.Config{})
	if err != nil {
		logger.Error("Failed to get db instance for migration", err)
	}

	s2 := "file://" + app.RootDir() + "/db/migrations"
	fmt.Println("Fff " + s2)
	m, err := migrate.NewWithDatabaseInstance(
		s2,
		"messages", driver)
	if err != nil {
		logger.Error("Failed to start db instance migration", err)
	}

	err = m.Migrate(s.config.CurrentVersion)
	if err == migrate.ErrNoChange {
		logger.Info("DB is up to date, no migration is needed")
	} else if err != nil {
		logger.Error("Failed to perform migration ", err)
		return err
	}

	logger.Info("DB is at migration version - ", s.config.CurrentVersion)
	return nil
}

func (s *Store) Close() {
	err := s.db.Close()
	if err != nil {
		common.GetLogger().Error("Unable to close db connection, you may leak resources")
	}
}
