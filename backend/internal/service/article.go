package service

import (
	"PersonalSite/backend/internal/model"
	"PersonalSite/backend/internal/repository"
)

type ArticleService struct {
	articleRepo  *repository.ArticleRepository
	categoryRepo *repository.CategoryRepository
	tagRepo      *repository.TagRepository
}

func NewArticleService(articleRepo *repository.ArticleRepository, categoryRepo *repository.CategoryRepository, tagRepo *repository.TagRepository) *ArticleService {
	return &ArticleService{
		articleRepo:  articleRepo,
		categoryRepo: categoryRepo,
		tagRepo:      tagRepo,
	}
}

type ArticleListResult struct {
	Total int64                  `json:"total"`
	Page  int                    `json:"page"`
	Size  int                    `json:"size"`
	List  []model.ArticleListResp `json:"list"`
}

func (s *ArticleService) List(page, size int, categoryID, tagID uint) (*ArticleListResult, error) {
	articles, total, err := s.articleRepo.List(page, size, categoryID, tagID)
	if err != nil {
		return nil, err
	}

	list := make([]model.ArticleListResp, 0, len(articles))
	for _, a := range articles {
		resp := model.ArticleListResp{
			ID:           a.ID,
			Title:        a.Title,
			Summary:      a.Summary,
			CategoryName: a.Category.Name,
			ViewCount:    a.ViewCount,
			CreatedAt:    a.CreatedAt,
		}
		for _, t := range a.Tags {
			resp.Tags = append(resp.Tags, t.Name)
		}
		list = append(list, resp)
	}

	return &ArticleListResult{
		Total: total,
		Page:  page,
		Size:  size,
		List:  list,
	}, nil
}

func (s *ArticleService) GetByID(id uint) (*model.ArticleDetailResp, error) {
	article, err := s.articleRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	// Increment view count
	_ = s.articleRepo.IncrementViewCount(id)

	resp := &model.ArticleDetailResp{
		ID:           article.ID,
		Title:        article.Title,
		Summary:      article.Summary,
		Content:      article.Content,
		CoverURL:     article.CoverURL,
		CategoryID:   article.CategoryID,
		CategoryName: article.Category.Name,
		ViewCount:    article.ViewCount + 1,
		CreatedAt:    article.CreatedAt,
		UpdatedAt:    article.UpdatedAt,
	}
	for _, t := range article.Tags {
		resp.Tags = append(resp.Tags, t.Name)
	}
	return resp, nil
}

type CreateArticleReq struct {
	Title      string `json:"title" binding:"required"`
	Summary    string `json:"summary"`
	Content    string `json:"content" binding:"required"`
	CoverURL   string `json:"cover_url"`
	CategoryID uint   `json:"category_id" binding:"required"`
	TagIDs     []uint `json:"tag_ids"`
}

func (s *ArticleService) Create(req *CreateArticleReq, userID uint) (*model.Article, error) {
	article := &model.Article{
		UserID:     userID,
		CategoryID: req.CategoryID,
		Title:      req.Title,
		Summary:    req.Summary,
		Content:    req.Content,
		CoverURL:   req.CoverURL,
	}

	if err := s.articleRepo.Create(article); err != nil {
		return nil, err
	}

	if len(req.TagIDs) > 0 {
		if err := s.articleRepo.ReplaceTags(article.ID, req.TagIDs); err != nil {
			return nil, err
		}
	}

	return article, nil
}

type UpdateArticleReq struct {
	Title      string `json:"title" binding:"required"`
	Summary    string `json:"summary"`
	Content    string `json:"content" binding:"required"`
	CoverURL   string `json:"cover_url"`
	CategoryID uint   `json:"category_id" binding:"required"`
	TagIDs     []uint `json:"tag_ids"`
}

func (s *ArticleService) Update(id uint, req *UpdateArticleReq) error {
	article, err := s.articleRepo.GetByID(id)
	if err != nil {
		return err
	}

	article.Title = req.Title
	article.Summary = req.Summary
	article.Content = req.Content
	article.CoverURL = req.CoverURL
	article.CategoryID = req.CategoryID

	if err := s.articleRepo.Update(article); err != nil {
		return err
	}

	if req.TagIDs != nil {
		return s.articleRepo.ReplaceTags(id, req.TagIDs)
	}

	return nil
}

func (s *ArticleService) Delete(id uint) error {
	return s.articleRepo.Delete(id)
}
