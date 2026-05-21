package models

import "time"

type Armor struct {
	ID      string          `json:"id" db:"id"`
	Name    string          `json:"name" db:"name"`
	Desc    string          `json:"desc" db:"desc"`
	Cost    float64         `json:"cost" db:"cost"`
	Attr    ArmorAttributes `json:"attr" db:"attr"`
	Visuals DataFile        `json:"visuals" db:"visuals"`
	Created time.Time       `json:"created" db:"created"`
	Updated time.Time       `json:"updated" db:"updated"`
}

type ArmorAttributes struct {
	Defense    int    `json:"defense"`
	Durability int    `json:"durability"`
	ArmorType  string `json:"armor_type"`
}
