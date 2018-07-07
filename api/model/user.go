package model

import (
	"time"
)

// User is Entity of user
type User struct {
	UserID    int64     `gorm:"primary_key;not null"`
	Name      string    `gorm:"size:255;not null"`
	Email     string    `gorm:"size:255;not null"`
	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`
	DeletedAt *time.Time
}
