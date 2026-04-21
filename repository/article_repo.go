package repository

import (
	"BlogPost/models"
	"database/sql"
)

type ArticleRepository struct {
	DB *sql.DB
}

func (r *ArticleRepository) CreateArticle(article models.Article) error {
	_, err := r.DB.Exec(
		"INSERT INTO articles (title, content, author_id) VALUES (?, ?, ?)",
		article.Title, article.Content, article.AuthorID,
	)
	return err
}

func (r *ArticleRepository) GetAllArticles() ([]models.Article, error) {
	rows, err := r.DB.Query("SELECT article_id, title, content, author_id FROM articles")
	if err != nil {
		return nil, err
	}

	var articles []models.Article

	for rows.Next() {
		var a models.Article
		rows.Scan(&a.ArticleID, &a.Title, &a.Content, &a.AuthorID)
		articles = append(articles, a)
	}

	return articles, nil
}
