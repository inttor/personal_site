package model

import "time"

type Tag struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"size:64;uniqueIndex;not null" json:"name"`
	CreatedAt time.Time `gorm:"precision:3" json:"created_at"`
	UpdatedAt time.Time `gorm:"precision:3" json:"updated_at"`
}
