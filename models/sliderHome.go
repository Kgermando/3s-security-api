package models

import "gorm.io/gorm"

type SliderHome struct {
	gorm.Model

	Title   string `gorm:"not null" json:"title"`
	SubTitle string `gorm:"not null" json:"sub_title"`
	Image  string `gorm:"not null" json:"image"`
	Signature string `json:"signature"`
}