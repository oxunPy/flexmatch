package models

import "time"

type Match struct {
	ID      string    `json:"id" db:"id"`
	Players []Player  `json:"players" db:"players"`
	Date    time.Time `json:"date" db:"date"`
	Type    MatchType `json:"type" db:"type"`
	Title   string    `json:"title" db:"title"`
}

type MatchPlayer struct {
	MatchID  string `json:"match_id" db:"match_id"`
	PlayerID int64  `json:"player_id" db:"player_id"`
}
