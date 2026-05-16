package main

import (
	"fmt"
	"log"
	"os"

	"PersonalSite/backend/config"
	"PersonalSite/backend/internal/model"
	"PersonalSite/backend/internal/router"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	cfg := config.Load()

	// Connect to MySQL
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Auto-migrate tables
	if err := db.AutoMigrate(
		&model.User{},
		&model.Category{},
		&model.Tag{},
		&model.Article{},
		&model.ArticleTag{},
	); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	// Seed default admin user if not exists
	seedAdmin(db)

	// Seed default categories if empty
	seedCategories(db)

	// Setup router
	r := router.Setup(db, cfg)

	// Start server
	addr := ":" + cfg.ServerPort
	log.Printf("Server starting on %s", addr)
	if err := r.Run(addr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
		os.Exit(1)
	}
}

func seedAdmin(db *gorm.DB) {
	var count int64
	db.Model(&model.User{}).Count(&count)
	if count > 0 {
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte("admin123"), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("Failed to hash password: %v", err)
	}

	admin := model.User{
		Username:     "admin",
		PasswordHash: string(hash),
		Email:        "admin@example.com",
	}

	if err := db.Create(&admin).Error; err != nil {
		log.Printf("Warning: failed to seed admin user: %v", err)
	} else {
		log.Println("Default admin user created (username: admin, password: admin123)")
	}
}

func seedCategories(db *gorm.DB) {
	var count int64
	db.Model(&model.Category{}).Count(&count)
	if count > 0 {
		return
	}

	categories := []model.Category{
		{Name: "技术", Description: "技术相关文章"},
		{Name: "生活", Description: "生活随笔"},
		{Name: "工具", Description: "工具使用与评测"},
	}

	for _, c := range categories {
		db.FirstOrCreate(&model.Category{}, model.Category{Name: c.Name})
	}
	log.Println("Default categories seeded")
}
