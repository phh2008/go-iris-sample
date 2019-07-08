package services

import (
	"com.phh/blog/models"
	"com.phh/blog/repositories"
)

type ArticleService interface {
	//查询文章
	Query(article models.Article) []models.Article
}

func NewArticleService(repo repositories.ArticleRepo) ArticleService {
	return &articleService{
		repo: repo,
	}
}

type articleService struct {
	repo repositories.ArticleRepo
}

func (a *articleService) Query(article models.Article) (results []models.Article) {
	return a.repo.Query(article)
}
