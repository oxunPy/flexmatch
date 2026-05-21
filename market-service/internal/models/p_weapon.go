package models

type PlayerWeapon struct {
	PlayerID int64  `json:"player_id" db:"player_id"`
	WeaponID string `json:"weapon_id" db:"weapon_id"`
}
