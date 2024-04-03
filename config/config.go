package config

import (
	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
	"hillel/logger"
	"hillel/models"
)

func InitConfig() (*models.Config, error) {
	if err := godotenv.Load(); err != nil {
		logger.Logger.Error(err)
		return nil, err
	}

	cfg := &models.Config{}
	if err := env.Parse(cfg); err != nil {
		logger.Logger.Error(err)
		return nil, err
	}

	return cfg, nil
}
