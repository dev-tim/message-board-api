package importer

import (
	"github.com/dev-tim/message-board-api/internal/app"
	"github.com/dev-tim/message-board-api/internal/app/store"
	"github.com/sirupsen/logrus"
)

type DataImporter struct {
	store  store.IStore
	logger *logrus.Logger
}

func New(store store.IStore, logger *logrus.Logger) *DataImporter {
	return &DataImporter{
		store:  store,
		logger: logger,
	}
}

func (s *DataImporter) Start() error {
	logger := s.logger

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
