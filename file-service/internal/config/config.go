package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	GrpcPort    int
	HttpPort    int
	DatabaseURL string
	StoragePath string
	BaseURL     string
}

func Load() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("err loading .env file: %w", err)
	}

	gport, err := strconv.Atoi(os.Getenv("GRPC_PORT"))
	if err != nil {
		return nil, fmt.Errorf("invalid http port config")
	}
	hport, err := strconv.Atoi(os.Getenv("HTTP_PORT"))
	if err != nil {
		return nil, fmt.Errorf("invalid grpc port config")
	}

	return &Config{
		GrpcPort:    gport,
		HttpPort:    hport,
		DatabaseURL: os.Getenv("DATABASE_URL"),
		StoragePath: os.Getenv("STORAGE_PATH"),
		BaseURL:     os.Getenv("BASE_URL"),
	}, nil
}
