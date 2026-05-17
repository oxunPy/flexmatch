package models

import "time"

type PlayerToken struct {
	ID        uint      `json:"id" db:"id"`
	Token     string    `json:"token" db:"token"`
	PlayerID  uint      `json:"player_id" db:"player_id"`
	Player    Player    `json:"player"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	ExpiredAt time.Time `json:"expired_at" db:"expired_at"`
}
