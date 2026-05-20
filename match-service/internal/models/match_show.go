package models

import "time"

type MatchShow struct {
	ID        string    `json:"id" db:"id"`
	MatchID   string    `json:"match_id" db:"match_id"`
	StartDate time.Time `json:"start_date" db:"start_date"`
	EndDate   time.Time `json:"end_date" db:"end_date"`
	Stream    chan []byte
}
