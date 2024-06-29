package config

import (
	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
	"github.com/opoccomaxao/installable-bot-concept/pkg/clients/db"
	"github.com/opoccomaxao/installable-bot-concept/pkg/server"
)

type Config struct {
	Server  server.Config `envPrefix:"SERVER_"`
	DBLocal db.Config     `envPrefix:"DB_LOCAL_"`
}

func Load() (*Config, error) {
	_ = godotenv.Load(".env")

	var cfg Config

	err := env.ParseWithOptions(&cfg, env.Options{
		UseFieldNameByDefault: false,
	})
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
