package config

import (
	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
	"github.com/labstack/gommon/log"
)

type Config struct {
	Port     int    `env:"PORT" envDefault:"1323"`
	LogLevel string `env:"LOG_LEVEL" envDefault:"info"`

	DBUser string `env:"DB_USER"`
	DBPass string `env:"DB_PASS"`
	DBHost string `env:"DB_HOST"`
	DBPort string `env:"DB_PORT"`
	DBName string `env:"DB_NAME"`
}

func MustLoad() *Config {
	if err := godotenv.Load(".env"); err != nil {
		log.Error(err)
	}

	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		log.Error(err)
	}

	return cfg
}
