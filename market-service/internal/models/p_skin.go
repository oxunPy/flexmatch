package models

type PlayerSkin struct {
	PlayerID int64  `json:"player_id" db:"player_id"`
	SkinID   string `json:"skin_id" db:"skin_id"`
}
