package common

import (
	"github.com/sirupsen/logrus"
	"sync"
)

type LoggerFactory struct {
	logger *logrus.Logger
	config *Config
}

var mutex = sync.Mutex{}
var loggerFactoryInstance *LoggerFactory

func NewLoggerFactory(config *Config) (*LoggerFactory, error) {
	if loggerFactoryInstance == nil {
		mutex.Lock()
		defer mutex.Unlock()

		if loggerFactoryInstance == nil {
			loggerFactoryInstance = &LoggerFactory{
				logger: logrus.New(),
				config: config,
			}

			if err := loggerFactoryInstance.configureLogger(); err != nil {
				return nil, err
			}
		}
	}

	return loggerFactoryInstance, nil
}

func GetLogger() *logrus.Logger {
	if loggerFactoryInstance == nil {
		logrus.Warn("Tried to get logger before initializing, panic")
		panic("Tried to get logger before initializing, panic")
	}

	return loggerFactoryInstance.logger
}

func (s *LoggerFactory) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}

	s.logger.SetLevel(level)
	return nil
}
