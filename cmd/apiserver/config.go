package main

import (
	"github.com/dev-tim/message-board-api/internal/app/apiserver"
	"github.com/dev-tim/message-board-api/internal/app/common"
	"github.com/dev-tim/message-board-api/internal/app/db/sqldb"
)

type Config struct {
	Api    *apiserver.Config
	Common *common.Config
	SqlDb  *sqldb.Config
}
