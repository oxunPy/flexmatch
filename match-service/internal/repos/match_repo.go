package repos

import (
	"context"
	"fmt"
	"match-service/internal/models"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func CreateMatch(pool *pgxpool.Pool, title string, date time.Time, typ models.MatchType) (*models.Match, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	var query = `
		INSERT INTO match(title, date, type)
		VALUES($1, $2, $3)
		RETURNING id, title, date, type
	`

	var match models.Match
	row := pool.QueryRow(ctx, query, title, date, typ)
	if err := row.Scan(
		&match.ID,
		&match.Title,
		&match.Date,
		&match.Type,
	); err != nil {
		return nil, err
	}

	return &match, nil
}

func DeleteMatch(pool *pgxpool.Pool, id string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	var query = `
		DELETE FROM match
		WHERE id = $1
	`
	commandTag, err := pool.Exec(ctx, query, id)
	if err != nil {
		return false, err
	}

	if commandTag.RowsAffected() == 0 {
		return false, fmt.Errorf("match with id %d not found", id)
	}

	return true, nil
}

func GetAllMatches(pool *pgxpool.Pool) ([]*models.Match, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	var query = `
		SELECT id, title, date, type
		FROM match
	`

	rows, err := pool.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	var matches []*models.Match
	for rows.Next() {
		var match models.Match
		if err := rows.Scan(
			&match.ID,
			&match.Title,
			&match.Date,
			&match.Type,
		); err != nil {
			continue
		}

		matches = append(matches, &match)
	}

	return matches, nil
}

func JoinMatch(pool *pgxpool.Pool, matchID string, playerID int64) (*models.MatchPlayer, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	var query = `
		INSERT INTO match_player(match_id, player_id) 
		VALUES($1, $2)
		RETURNING match_id, player_id
	`

	var mp models.MatchPlayer
	row := pool.QueryRow(ctx, query, matchID, playerID)
	if err := row.Scan(
		&mp.MatchID,
		&mp.PlayerID,
	); err != nil {
		return nil, err
	}

	return &mp, nil
}
