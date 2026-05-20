package database

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Connect(url string) (*pgxpool.Pool, error) {
	ctx := context.Background()
	config, err := pgxpool.ParseConfig(url)
	if err != nil {
		log.Printf("unable to parse database url: %v\n", err)
		return nil, err
	}

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		log.Printf("unable to create connection pool: %v\n", err)
		return nil, err
	}

	err = pool.Ping(ctx)
	if err != nil {
		log.Printf("unable to ping database: %v\n", err)
		pool.Close()
		return nil, err
	}

	log.Println("successfully connected to PostgreSQL database")
	return pool, nil
}
