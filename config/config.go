package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Host       string `env-required:"true" yaml:"host"`
	Port       string `env-required:"true" yaml:"port"`
	User       string `env-required:"true" yaml:"user"`
	Password   string `env:"DB_PASSWORD"`
	DBName     string `env-required:"true" yaml:"dbname"`
	SSLMode    string `env-required:"true" yaml:"sslmode"`
	ServerPort string `env-required:"true" yaml:"serverPort"`
}

func NewConfig() (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadConfig("./config/config.yml", cfg)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
