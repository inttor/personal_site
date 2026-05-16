package router

import (
	"PersonalSite/backend/config"
	"PersonalSite/backend/internal/handler"
	"PersonalSite/backend/internal/middleware"
	"PersonalSite/backend/internal/repository"
	"PersonalSite/backend/internal/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Setup(db *gorm.DB, cfg *config.Config) *gin.Engine {
	r := gin.Default()

	// CORS
	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	// Static files for uploaded images
	r.Static("/uploads", cfg.UploadDir)

	// Repositories
	userRepo := repository.NewUserRepository(db)
	articleRepo := repository.NewArticleRepository(db)
	categoryRepo := repository.NewCategoryRepository(db)
	tagRepo := repository.NewTagRepository(db)

	// Services
	articleSvc := service.NewArticleService(articleRepo, categoryRepo, tagRepo)

	// Handlers
	authHandler := handler.NewAuthHandler(userRepo, cfg)
	articleHandler := handler.NewArticleHandler(articleSvc)
	categoryHandler := handler.NewCategoryHandler(categoryRepo)
	tagHandler := handler.NewTagHandler(tagRepo)
	uploadHandler := handler.NewUploadHandler(cfg)

	// Routes
	api := r.Group("/api/v1")
	{
		// Public
		api.POST("/auth/login", authHandler.Login)

		api.GET("/articles", articleHandler.List)
		api.GET("/articles/:id", articleHandler.GetByID)

		api.GET("/categories", categoryHandler.List)
		api.GET("/tags", tagHandler.List)

		// Auth required
		auth := api.Group("")
		auth.Use(middleware.Auth(cfg))
		{
			auth.POST("/articles", articleHandler.Create)
			auth.PUT("/articles/:id", articleHandler.Update)
			auth.DELETE("/articles/:id", articleHandler.Delete)
			auth.POST("/upload", uploadHandler.Upload)
		}
	}

	return r
}
