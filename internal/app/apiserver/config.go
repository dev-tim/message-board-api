package apiserver

import (
	"github.com/dev-tim/message-board-api/internal/app/common"
	"github.com/dev-tim/message-board-api/internal/app/store"
)

type Config struct {
	BindAddress string `toml:"bind_address"`
	Store       *store.Config
	Common      *common.Config
}

func NewConfig() *Config {
	return &Config{
		Store:  store.NewConfig(),
		Common: common.NewLoggerConfig(),
	}
}
