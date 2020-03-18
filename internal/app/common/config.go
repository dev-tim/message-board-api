package common

type Config struct {
	LogLevel string `toml:"log_level"`
}

func NewLoggerConfig() *Config {
	return &Config{
		LogLevel: "Debug",
	}
}
