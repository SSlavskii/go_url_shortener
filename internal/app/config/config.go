package config

import (
	"github.com/caarlos0/env/v6"
)

type Config struct {
	ServerAdress string `env:"SERVER_ADDRESS" envDefault:":8080"`
	BaseURL      string `env:"BASE_URL" envDefault:"http://localhost:8080"`
}

func New() (*Config, error) {
	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		return &cfg, err
	}
	return &cfg, nil
}
