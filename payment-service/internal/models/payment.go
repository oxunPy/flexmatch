package models

import "time"

type PaymentType int

const (
	IN PaymentType = iota
	OUT
)

type Payment struct {
	ID       int64       `json:"id" db:"id"`
	ItemID   string      `json:"item_id" db:"item_id"`
	PlayerID int64       `json:"player_id" db:"player_id"`
	Type     PaymentType `json:"type" db:"type"`
	Amount   float64     `json:"amount" db:"amount"`
	WalletID int64       `json:"wallet_id" db:"wallet_id"`
	Created  time.Time   `json:"created" db:"created"`
}
