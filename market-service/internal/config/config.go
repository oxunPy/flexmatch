package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL     string
	HttpPort        int
	GrpcPort        int
	FileGrpcPort    int
	PaymentGrpcPort int
}

func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("error loading .env config file")
	}

	hport, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		return nil, fmt.Errorf("err config PORT")
	}

	gport, err := strconv.Atoi(os.Getenv("GPORT"))
	if err != nil {
		return nil, fmt.Errorf("err config GPORT")
	}

	fileGport, err := strconv.Atoi(os.Getenv("FILE_SERVICE_GPORT"))
	if err != nil {
		return nil, fmt.Errorf("err config FILE_SERVICE_GPORT")
	}

	paymentGport, err := strconv.Atoi(os.Getenv("PAYMENT_SERVICE_GPORT"))
	if err != nil {
		return nil, fmt.Errorf("err config PAYMENT_SERVICE_GPORT")
	}

	return &Config{
		DatabaseURL:     os.Getenv("DATABASE_URL"),
		HttpPort:        hport,
		GrpcPort:        gport,
		FileGrpcPort:    fileGport,
		PaymentGrpcPort: paymentGport,
	}, nil
}
