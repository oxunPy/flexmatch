package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL string
	Port        int
	GPort       int
	JWTSecret   string
}

func Load() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found, using environment variables")
		return nil, err
	}

	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		log.Println("Warning: invalid port property")
		return nil, err
	}

	gport, err := strconv.Atoi(os.Getenv("GRPC_PORT"))
	if err != nil {
		log.Println("Warning: invalid gport property")
		return nil, err
	}

	return &Config{
		DatabaseURL: os.Getenv("DATABASE_URL"),
		Port:        port,
		GPort:       gport,
		JWTSecret:   os.Getenv("JWT_SECRET"),
	}, nil
}
