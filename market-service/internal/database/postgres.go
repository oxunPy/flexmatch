package database

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Connect(url string) (*pgxpool.Pool, error) {
	ctx := context.Background()
	cfg, err := pgxpool.ParseConfig(url)
	if err != nil {
		return nil, fmt.Errorf("error parse database url configuration: %w", err)
	}

	pool, err := pgxpool.NewWithConfig(ctx, cfg)
	if err != nil {
		return nil, fmt.Errorf("error connect to database: %w", err)
	}

	if err := pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("error ping database: %w", err)
	}

	slog.Info("successfully connected to the PostgreSQL database!")
	return pool, nil
}
