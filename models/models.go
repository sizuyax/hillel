package models

type Config struct {
	Port     string `env:"PORT"`
	LogLevel string `env:"LOG_LEVEL"`
}
