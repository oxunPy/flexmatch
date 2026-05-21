package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL string
	Port        string
}

func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("error loading .env config file")
	}

	return &Config{
		DatabaseURL: os.Getenv("DATABASE_URL"),
		Port:        os.Getenv("PORT"),
	}, nil
}
