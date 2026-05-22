package repos

import (
	"auth-service/internal/database"
	"auth-service/internal/models"
	"context"
	"time"
)

type PlayerRepo struct {
	storage *database.PostgresStorage
}

func NewPlayerRepo(storage *database.PostgresStorage) *PlayerRepo {
	return &PlayerRepo{
		storage: storage,
	}
}

func (r *PlayerRepo) CreateNewPlayer(player models.Player) (*models.Player, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	var query = `
		INSERT INTO players(username, firstname, lastname, email, password)
		VALUES($1, $2, $3, $4, $5)
		RETURNING id, username, firstname, lastname, email, password, disabled, created_at, updated_at
	`

	var p models.Player
	err := r.storage.QueryRow(ctx, query,
		player.Username,
		player.Firstname,
		player.Lastname,
		player.Email,
		player.Password,
	).Scan(
		&p.ID,
		&p.Username,
		&p.Firstname,
		&p.Lastname,
		&p.Email,
		&p.Password,
		&p.Disabled,
		&p.CreatedAt,
		&p.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &p, nil
}

func (r *PlayerRepo) GetPlayer(username string) (*models.Player, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	var query = `
		SELECT id, username, firstname, lastname, email, password, disabled, created_at, updated_at
		FROM players
		WHERE username = $1 or email = $1
	`

	var p models.Player
	err := r.storage.QueryRow(ctx, query, username).Scan(
		&p.ID,
		&p.Username,
		&p.Firstname,
		&p.Lastname,
		&p.Email,
		&p.Password,
		&p.Disabled,
		&p.CreatedAt,
		&p.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &p, nil
}
