package routes

import (
	controller "BlogPost/controllers"
	"BlogPost/middleware"
	"BlogPost/repository"
	"BlogPost/service"
	"database/sql"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, db *sql.DB) {

	userRepo := &repository.UserRepository{DB: db}
	articleRepo := &repository.ArticleRepository{DB: db}

	authService := &service.AuthService{UserRepo: userRepo}
	articleService := &service.ArticleService{ArticleRepo: articleRepo}

	authController := &controller.AuthController{Service: authService}
	articleController := &controller.ArticleController{Service: articleService}

	auth := r.Group("/")
	auth.Use(middleware.AuthMiddleware())

	r.POST("/signup", authController.Signup)
	r.POST("/login", authController.Login)

	r.POST("/articles", articleController.CreateArticle)
	r.GET("/articles", articleController.GetArticles)
}
