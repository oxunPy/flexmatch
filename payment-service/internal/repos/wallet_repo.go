package repos

import (
	"context"
	"payment-service/internal/models"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func CreateWallet(pool *pgxpool.Pool, playerID int64) (*models.Wallet, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	var query = `
		INSERT INTO wallets(player_id, balance)
		VALUES($1, $2)
		RETURNING id, player_id, balance, created
	`

	var wallet models.Wallet
	err := pool.
		QueryRow(ctx, query, playerID, float64(0)).
		Scan(&wallet.ID, &wallet.PlayerID, &wallet.Balance, &wallet.Created)

	if err != nil {
		return nil, err
	}

	return &wallet, err
}

func GetMyWallets(pool *pgxpool.Pool, playerID int64) ([]*models.Wallet, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	var query = `
		SELECT id, player_id, balance, created
		FROM wallets
		WHERE player_id = $1
	`

	rows, err := pool.Query(ctx, query, playerID)
	if err != nil {
		return nil, err
	}

	var wallets = make([]*models.Wallet, 0)
	for rows.Next() {
		var wlt models.Wallet
		if err := rows.Scan(&wlt.ID, &wlt.PlayerID, &wlt.Balance, &wlt.Created); err != nil {
			return nil, err
		}

		wallets = append(wallets, &wlt)
	}

	return wallets, nil
}

func GetWalletsList(pool *pgxpool.Pool) ([]*models.Wallet, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	var query = `
		SELECT id, player_id, balance, created
		FROM wallets
	`

	rows, err := pool.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	var wallets = make([]*models.Wallet, 0)
	for rows.Next() {
		var wlt models.Wallet
		if err := rows.Scan(&wlt.ID, &wlt.PlayerID, &wlt.Balance, &wlt.Created); err != nil {
			return nil, err
		}
		wallets = append(wallets, &wlt)
	}

	return wallets, nil
}
