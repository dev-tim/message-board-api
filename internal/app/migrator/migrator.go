package migrator

import (
	"github.com/dev-tim/message-board-api/internal/app"
	"github.com/dev-tim/message-board-api/internal/app/common"
	"github.com/dev-tim/message-board-api/internal/app/store"
)

type DataImporter struct {
	name   string
	config *Config
	store  *store.Store
}

func New(config *Config) *DataImporter {
	return &DataImporter{
		name:   "DataImporter",
		config: config,
	}
}

func (s *DataImporter) Start() error {
	if _, err := common.NewLoggerFactory(s.config.Common); err != nil {
		return err
	}
	logger := common.GetLogger()

	if err := s.configureStore(); err != nil {
		return err
	}

	logger.Info("Started data importer")
	messages, err := ReadCSVFromFile(app.RootDir() + "/messages.csv")
	if err != nil {
		logger.Error("Failed from csv messages", err)
	}

	for _, m := range messages {
		_, err := s.store.Messages().Create(m)
		if err != nil {
			logger.Error("Failed to store message", m.Id)
		}
	}

	return nil
}

func (s *DataImporter) configureStore() error {
	st := store.New(s.config.Store)
	if err := st.Open(); err != nil {
		return err
	}

	s.store = st
	return nil
}
