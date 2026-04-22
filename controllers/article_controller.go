package controller

import (
	"BlogPost/config"
	"BlogPost/models"
	"BlogPost/service"
	"encoding/json"

	"github.com/gin-gonic/gin"
)

type ArticleController struct {
	Service *service.ArticleService
}

func (c *ArticleController) CreateArticle(ctx *gin.Context) {
	var input struct {
		Title   string
		Content string
	}

	ctx.BindJSON(&input)

	userID, _ := ctx.Get("user_id")
	article := models.Article{
		Title:    input.Title,
		Content:  input.Content,
		AuthorID: userID.(int64),
	}
	err := c.Service.CreateArticle(article)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": "Created"})
}

func (c *ArticleController) GetArticles(ctx *gin.Context) ([]models.Article, error) {

	cacheKey := "articles_feed"
	val, err := config.RDB.Get(config.Ctx, cacheKey).Result()

	if err == nil {
		var articles []models.Article
		json.Unmarshal([]byte(val), &articles)
		return articles, nil
	}
	articles, err := c.Service.GetArticles()
	if err != nil {
		return nil, err
	}
	data, _ := json.Marshal(articles)

	config.RDB.Set(config.Ctx, cacheKey, data, 0)

	return articles, nil
}
