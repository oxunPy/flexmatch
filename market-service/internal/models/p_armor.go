package models

type PlayerArmor struct {
	PlayerID int64  `json:"player_id" db:"player_id"`
	ArmorID  string `json:"armor_id" db:"armor_id"`
}
