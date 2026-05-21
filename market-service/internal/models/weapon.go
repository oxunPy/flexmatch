package models

import "time"

type WeaponType int

const (
	Sword WeaponType = iota
	Rifle
	Bow
)

type Weapon struct {
	ID      string           `json:"id" db:"id"`
	Name    string           `json:"name" db:"name"`
	Desc    string           `json:"desc" db:"desc"`
	Type    WeaponType       `json:"type" db:"weapon_type"`
	Cost    float64          `json:"cost" db:"cost"`
	Attr    WeaponAttributes `json:"attr" db:"attr"`
	Visuals DataFile         `json:"visuals" db:"visuals"`
	Created time.Time        `json:"created" db:"created"`
	Updated time.Time        `json:"updated" db:"updated"`
}

type WeaponAttributes struct {
	Damage      int     `json:"damage"`
	AttackSpeed float64 `json:"attack_speed"`
	WeaponType  string  `json:"weapon_type"`
}
