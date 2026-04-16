package models

import "time"

type User struct {
	ID        uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Username  string `gorm:"uniqueIndex;not null" json:"username"`
	Password  string `gorm:"not null" json:"-"`
	Role      string `gorm:"not null" json:"role"`
	CreatedAt time.Time `json:"created_at"`
}