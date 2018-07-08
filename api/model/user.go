package model

import (
	"time"
)

// User is Entity of user
type User struct {
	UserID    int64      `gorm:"primary_key;not null" json:"userId"`
	Name      string     `gorm:"size:255;not null" json:"name"`
	Email     string     `gorm:"size:255;not null" json:"email"`
	CreatedAt time.Time  `gorm:"not null" json:"createdAt"`
	UpdatedAt time.Time  `gorm:"not null" json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
}
