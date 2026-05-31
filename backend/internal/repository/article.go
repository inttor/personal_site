package repository

import (
	"PersonalSite/backend/internal/model"

	"gorm.io/gorm"
)

type ArticleRepository struct {
	db *gorm.DB
}

func NewArticleRepository(db *gorm.DB) *ArticleRepository {
	return &ArticleRepository{db: db}
}

func (r *ArticleRepository) List(page, size int, categoryID, tagID uint) ([]model.Article, int64, error) {
	var articles []model.Article
	var total int64

	query := r.db.Model(&model.Article{}).Preload("Category").Preload("Tags")

	if categoryID > 0 {
		query = query.Where("category_id = ?", categoryID)
	}
	if tagID > 0 {
		query = query.Where("id IN (SELECT article_id FROM article_tags WHERE tag_id = ?)", tagID)
	}

	query.Count(&total)

	err := query.Scopes(func(db *gorm.DB) *gorm.DB {
		return db.Offset((page - 1) * size).Limit(size)
	}).Order("created_at DESC").Find(&articles).Error

	return articles, total, err
}

func (r *ArticleRepository) GetByID(id uint) (*model.Article, error) {
	var article model.Article
	err := r.db.Preload("Category").Preload("Tags").First(&article, id).Error
	if err != nil {
		return nil, err
	}
	return &article, nil
}

func (r *ArticleRepository) Create(article *model.Article) error {
	return r.db.Create(article).Error
}

func (r *ArticleRepository) Update(article *model.Article) error {
	return r.db.Model(&model.Article{}).Where("id = ?", article.ID).Updates(map[string]interface{}{
		"title":       article.Title,
		"summary":     article.Summary,
		"content":     article.Content,
		"cover_url":   article.CoverURL,
		"category_id": article.CategoryID,
	}).Error
}

func (r *ArticleRepository) Delete(id uint) error {
	return r.db.Delete(&model.Article{}, id).Error
}

func (r *ArticleRepository) IncrementViewCount(id uint) error {
	return r.db.Model(&model.Article{}).Where("id = ?", id).UpdateColumn("view_count", gorm.Expr("view_count + 1")).Error
}

func (r *ArticleRepository) ReplaceTags(articleID uint, tagIDs []uint) error {
	tx := r.db.Begin()
	if err := tx.Where("article_id = ?", articleID).Delete(&model.ArticleTag{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	for _, tagID := range tagIDs {
		if err := tx.Create(&model.ArticleTag{ArticleID: articleID, TagID: tagID}).Error; err != nil {
			tx.Rollback()
			return err
		}
	}
	return tx.Commit().Error
}
