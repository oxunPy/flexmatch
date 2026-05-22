package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Connect(url string) (*pgxpool.Pool, error) {
	ctx := context.Background()
	cfg, err := pgxpool.ParseConfig(url)
	if err != nil {
		return nil, fmt.Errorf("err parsing database url config")
	}

	pool, err := pgxpool.NewWithConfig(ctx, cfg)
	if err != nil {
		return nil, fmt.Errorf("err connect to database")
	}

	if err := pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("err ping database")
	}

	return pool, nil
}
