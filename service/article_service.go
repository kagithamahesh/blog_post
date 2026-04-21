package service

import (
	"BlogPost/models"
	"BlogPost/repository"
)

type ArticleService struct {
	ArticleRepo *repository.ArticleRepository
}

func (s *ArticleService) CreateArticle(article models.Article) error {
	return s.ArticleRepo.CreateArticle(article)
}

func (s *ArticleService) GetArticles() ([]models.Article, error) {
	return s.ArticleRepo.GetAllArticles()
}
