package models

import "gorm.io/gorm"

type Blog struct {
	gorm.Model

	Title     string   `gorm:"not null" json:"title"`
	Content   string   `gorm:"not null" json:"content"`
	Images    []string `gorm:"type:text[]" json:"images"`
	Author    string   `gorm:"not null" json:"author"`
	Tags      []string `gorm:"type:text[]" json:"tags"`
	Status    bool     `gorm:"not null" json:"status"`
	Signature string   `json:"signature"`
}
