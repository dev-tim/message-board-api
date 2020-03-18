package main

import (
	"github.com/dev-tim/message-board-api/internal/app/common"
	"github.com/dev-tim/message-board-api/internal/app/db/sqldb"
)

type Config struct {
	SqlDb  *sqldb.Config
	Common *common.Config
}
