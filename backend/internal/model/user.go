package model

import "time"

type User struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	Username     string    `gorm:"size:64;uniqueIndex;not null" json:"username"`
	PasswordHash string    `gorm:"size:255;not null" json:"-"`
	Email        string    `gorm:"size:128" json:"email"`
	CreatedAt    time.Time `gorm:"precision:3" json:"created_at"`
	UpdatedAt    time.Time `gorm:"precision:3" json:"updated_at"`
}
