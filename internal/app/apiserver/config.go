package apiserver

type Config struct {
	BindAddress string `toml:"bind_address"`
}

func NewConfig() *Config {
	return &Config{}
}
