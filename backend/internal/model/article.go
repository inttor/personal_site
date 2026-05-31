package model

import (
	"time"

	"gorm.io/gorm"
)

type Article struct {
	ID         uint           `gorm:"primaryKey" json:"id"`
	UserID     uint           `gorm:"not null" json:"user_id"`
	CategoryID uint           `gorm:"not null" json:"category_id"`
	Title      string         `gorm:"size:255;not null" json:"title"`
	Summary    string         `gorm:"size:512" json:"summary"`
	Content    string         `gorm:"type:longtext;not null" json:"content"`
	CoverURL   string         `gorm:"size:512" json:"cover_url"`
	ViewCount  uint           `gorm:"default:0" json:"view_count"`
	CreatedAt  time.Time      `gorm:"precision:3" json:"created_at"`
	UpdatedAt  time.Time      `gorm:"precision:3" json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"precision:3" json:"-"`

	// Associations
	Category Category `gorm:"foreignKey:CategoryID" json:"category,omitempty"`
	Tags     []Tag    `gorm:"many2many:article_tags;joinForeignKey:ArticleID;joinReferences:TagID" json:"tags,omitempty"`
}

// ArticleListResp is the response struct for article list items
type ArticleListResp struct {
	ID           uint      `json:"id"`
	Title        string    `json:"title"`
	Summary      string    `json:"summary"`
	CategoryName string    `json:"category_name"`
	Tags         []string  `json:"tags"`
	ViewCount    uint      `json:"view_count"`
	CreatedAt    time.Time `json:"created_at"`
}

// ArticleDetailResp is the response struct for article detail
type ArticleDetailResp struct {
	ID           uint      `json:"id"`
	Title        string    `json:"title"`
	Summary      string    `json:"summary"`
	Content      string    `json:"content"`
	CoverURL     string    `json:"cover_url"`
	CategoryID   uint      `json:"category_id"`
	CategoryName string    `json:"category_name"`
	Tags         []string  `json:"tags"`
	ViewCount    uint      `json:"view_count"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
