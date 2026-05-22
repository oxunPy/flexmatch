package database

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresStorage struct {
	*pgxpool.Pool
}

func NewStorage(databaseUrl string) (*PostgresStorage, error) {
	ctx := context.Background()
	config, err := pgxpool.ParseConfig(databaseUrl)
	if err != nil {
		log.Printf("Unable to parse DATABASE_URL: %v\n", err)
	}

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		log.Printf("Unable to create connection pool: %v\n", err)
	}

	err = pool.Ping(ctx)
	if err != nil {
		log.Printf("Unable to ping database: %v\n", err)
		pool.Close()
		return nil, err
	}

	log.Println("Successfully connected to PostgreSQL database")
	return &PostgresStorage{
		Pool: pool,
	}, nil
}
