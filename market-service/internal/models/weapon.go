package models

import "time"

type WeaponType int

const (
	Sword WeaponType = iota
	Rifle
	Bow
)

type Weapon struct {
	ID      string     `json:"id" db:"id"`
	Name    string     `json:"name" db:"name"`
	Desc    string     `json:"desc" db:"desc"`
	Type    WeaponType `json:"type" db:"type"`
	Cost    float64    `json:"cost" db:"cost"`
	Visuals DataFile   `json:"visuals" db:"visuals"`
	Created time.Time  `json:"created" db:"updated"`
	Updated time.Time  `json:"updated" db:"updated"`
}

type WeaponAttributes struct {
	Damage      int     `json:"damage"`
	AttackSpeed float64 `json:"attack_speed"`
	WeaponType  string  `json:"weapon_type"`
}
