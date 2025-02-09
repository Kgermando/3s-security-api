package models

import "gorm.io/gorm"

type Team struct {
	gorm.Model

	Name      string `gorm:"not null" json:"name"`
	Position  string `gorm:"not null" json:"position"` // Position occuper dans l'entreprise
	Avatar    string `gorm:"not null" json:"avatar"`
	Facebook string `json:"facebook"`
	LinkedIn string `json:"linkedIn"`
	Twitter string `json:"twitter"`
	Signature string `json:"signature"`
}

