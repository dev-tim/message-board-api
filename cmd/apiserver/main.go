package main

import (
	"database/sql"
	"flag"
	"github.com/BurntSushi/toml"
	"github.com/dev-tim/message-board-api/internal/app/apiserver"
	"github.com/dev-tim/message-board-api/internal/app/common"
	"github.com/dev-tim/message-board-api/internal/app/db/sqldb"
	"github.com/dev-tim/message-board-api/internal/app/store/sqlstore"
	"log"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "Api server config file path")
}

func main() {
	flag.Parse()

	config := &Config{}
	if _, err := toml.DecodeFile(configPath, config); err != nil {
		log.Fatal(err)
	}

	_, err := common.NewLoggerFactory(config.Common)
	if err != nil {
		log.Fatal("Unable to init logger", err)
	}

	db, err := ProvideDB(err, config)
	if err != nil {
		log.Fatal("Unable to open db", err)
	}
	defer db.Close()

	store := sqlstore.New(db, common.GetLogger())

	server := apiserver.New(config.Api, store, common.GetLogger())
	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}

func ProvideDB(err error, config *Config) (*sql.DB, error) {
	db, err := sqldb.Open(config.SqlDb)
	if err != nil {
		return nil, err
	}

	if err := sqldb.Migrate(db, config.SqlDb, common.GetLogger()); err != nil {
		return nil, err
	}

	return db, nil
}
