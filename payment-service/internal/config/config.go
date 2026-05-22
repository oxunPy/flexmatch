package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL string
	Port        int
	GPort       int
}

func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("err loading .env file")
	}

	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		return nil, fmt.Errorf("err configuration: PORT")
	}

	gport, err := strconv.Atoi(os.Getenv("GRPC_PORT"))
	if err != nil {
		return nil, fmt.Errorf("err configuration: GRPC_PORT")
	}

	return &Config{
		DatabaseURL: os.Getenv("DATABASE_URL"),
		Port:        port,
		GPort:       gport,
	}, nil
}
