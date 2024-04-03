package models

type Config struct {
	Port     string `env:"PORT" envDefault:"8081"`
	LogLevel string `env:"LOG_LEVEL" envDefault:"info"`
}
