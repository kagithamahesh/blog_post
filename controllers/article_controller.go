package controller

import (
	"BlogPost/models"
	"BlogPost/service"

	"github.com/gin-gonic/gin"
)

type ArticleController struct {
	Service *service.ArticleService
}

func (c *ArticleController) CreateArticle(ctx *gin.Context) {
	var article models.Article
	ctx.BindJSON(&article)

	err := c.Service.CreateArticle(article)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": "Created"})
}

func (c *ArticleController) GetArticles(ctx *gin.Context) {
	articles, _ := c.Service.GetArticles()
	ctx.JSON(200, articles)
}
