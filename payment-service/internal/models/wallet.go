package models

import "time"

type Wallet struct {
	ID       int64     `json:"id" db:"id"`
	PlayerID int64     `json:"player_id" db:"player_id"`
	Balance  int64     `json:"balance" db:"balance"`
	Created  time.Time `json:"created"`
}
