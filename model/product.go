package model

import "gorm.io/gorm"

// Product struct
type Product struct {
	gorm.Model
	Product     string `gorm:"not null" json:"id"`
	Description string `gorm:"not null" json:"description"`
	Amount      int    `gorm:"not null" json:"amount"`
}
