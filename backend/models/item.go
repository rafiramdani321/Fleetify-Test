package models

import "time"

type Item struct {
	ID        uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Code      string `gorm:"uniqueIndex;not null" json:"code"`
	Name      string `gorm:"not null" json:"name"`
	Price     int64  `gorm:"not null" json:"price"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}