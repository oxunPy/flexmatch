package repos

import (
	"auth-service/internal/models"
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func CreatePlayerToken(pool *pgxpool.Pool, pt models.PlayerToken) (*models.PlayerToken, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	query := `
		INSERT INTO player_tokens(token, player_id, expired_at)
		VALUES($1, $2, $3)
		RETURNING id, token, player_id, created_at, expired_at
	`

	var newPt models.PlayerToken
	err := pool.QueryRow(ctx, query, pt.Token, pt.ID, pt.ExpiredAt).Scan(
		&newPt.ID,
		&newPt.Token,
		&newPt.PlayerID,
		&newPt.CreatedAt,
		&newPt.ExpiredAt,
	)
	if err != nil {
		return nil, err
	}
	newPt.Player = pt.Player

	return &newPt, nil
}
