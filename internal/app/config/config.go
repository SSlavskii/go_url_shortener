package config

import (
	"flag"

	"github.com/caarlos0/env/v6"
)

type Config struct {
	ServerAdress    string `env:"SERVER_ADDRESS" envDefault:":8080"`
	BaseURL         string `env:"BASE_URL" envDefault:"http://localhost:8080"`
	FileStoragePath string `env:"FILE_STORAGE_PATH" envDefault:"storage.txt"`
}

func New() (*Config, error) {
	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		return &cfg, err
	}
	readFlags(&cfg)
	return &cfg, nil
}

func readFlags(cfg *Config) {
	flag.StringVar(&cfg.ServerAdress, "a", cfg.ServerAdress, "SERVER_ADDRESS")
	flag.StringVar(&cfg.BaseURL, "b", cfg.BaseURL, "BASE_URL")
	flag.StringVar(&cfg.FileStoragePath, "f", cfg.FileStoragePath, "FILE_STORAGE_PATH")
	flag.Parse()
}
