package sqldb

type Config struct {
	DbUrl                   string `toml:"db_url"`
	MaxIdleConnections      int    `toml:"max_idle_connections"`
	MaxOpenConnections      int    `toml:"max_open_connections"`
	AquireConnectionTimeout int    `toml:"aquire_connection_timeout_seconds"`
	CurrentVersion          uint   `toml:"current_version"`
}

func NewConfig() *Config {
	return &Config{}
}
