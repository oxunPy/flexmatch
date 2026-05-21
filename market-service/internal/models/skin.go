package models

import "time"

type Skin struct {
	ID      string         `json:"id" db:"id"`
	Name    string         `json:"name" db:"name"`
	Visuals DataFile       `json:"visuals" db:"visuals"`
	Attr    SkinAttributes `json:"attr" db:"attr"`
	Cost    float64        `json:"cost" db:"cost"`
	Created time.Time      `json:"created" db:"created"`
	Updated time.Time      `json:"updated" db:"updated"`
}

type SkinAttributes struct {
	Effects DataFile `json:"effects" db:"effects"`
}
