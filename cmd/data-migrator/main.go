package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	"github.com/dev-tim/message-board-api/internal/app/migrator"
	"log"
	"os"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/migrator.toml", "Data importer config path")
}

func main() {
	flag.Parse()

	config := migrator.NewConfig()
	if _, err := toml.DecodeFile(configPath, config); err != nil {
		log.Fatal(err)
	}

	importer := migrator.New(config)
	if err := importer.Start(); err != nil {
		log.Fatal(err)
	}

	os.Exit(0)
}
