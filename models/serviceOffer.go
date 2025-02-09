package models

import "gorm.io/gorm"

type ServiceOffer struct  {
	gorm.Model

	Title     string   `gorm:"not null" json:"title"`
	Content   string   `gorm:"not null" json:"content"`
	Images    []string `gorm:"type:text[]" json:"images"`
	Signature string   `json:"signature"`
}
