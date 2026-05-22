package repos

import (
	"context"
	"payment-service/internal/database"
	"payment-service/internal/models"
	"time"
)

type PaymentRepo struct {
	storage *database.PostgresStorage
}

func NewPaymentRepo(storage *database.PostgresStorage) *PaymentRepo {
	return &PaymentRepo{
		storage: storage,
	}
}

func (r *PaymentRepo) CreatePayment(pay models.Payment) (*models.Payment, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	var query = `
		INSERT INTO payments(item_id, player_id, type, amount, wallet_id)
		VALUES($1::uuid, $2, $3, $4, $5)
		RETURNING id, item_id, player_id, wallet_id, type, amount, created
	`

	var newPay models.Payment
	err := r.storage.QueryRow(ctx, query,
		pay.ItemID,
		pay.PlayerID,
		pay.Type,
		pay.Amount,
		pay.WalletID,
	).Scan(&newPay.ID, &newPay.ItemID, &newPay.PlayerID, &newPay.WalletID, &newPay.Type, &newPay.Amount, &newPay.Created)
	if err != nil {
		return nil, err
	}

	return &newPay, nil
}

func (r *PaymentRepo) GetAllPaymentList() ([]*models.Payment, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	var query = `
		SELECT id, item_id, player_id, wallet_id, type, amount, created
		FROM payments
	`

	rows, err := r.storage.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	list := make([]*models.Payment, 0)
	for rows.Next() {
		var pay models.Payment
		err := rows.Scan(&pay.ID, &pay.ItemID, &pay.PlayerID, &pay.WalletID, &pay.Type, &pay.Amount, &pay.Created)
		if err != nil {
			continue
		}

		list = append(list, &pay)
	}

	return list, nil
}
