package config

type Config struct {
	Addr     string
	LogLevel string
}

func New() *Config {
	return &Config{}
}
