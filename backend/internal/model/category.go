package model

import "time"

type Category struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Name        string    `gorm:"size:64;uniqueIndex;not null" json:"name"`
	Description string    `gorm:"size:255" json:"description"`
	CreatedAt   time.Time `gorm:"precision:3" json:"created_at"`
	UpdatedAt   time.Time `gorm:"precision:3" json:"updated_at"`
}
