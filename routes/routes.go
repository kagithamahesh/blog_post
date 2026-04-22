package routes

import (
	"fmt"
	controller "BlogPost/controllers"
	"BlogPost/middleware"
	"BlogPost/repository"
	"BlogPost/service"
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, db *sql.DB) error {
	// Validate database connection
	if db == nil {
		return fmt.Errorf("database connection is nil")
	}

	// Initialize repositories
	userRepo := &repository.UserRepository{DB: db}
	articleRepo := &repository.ArticleRepository{DB: db}

	// Initialize services
	authService := &service.AuthService{UserRepo: userRepo}
	articleService := &service.ArticleService{ArticleRepo: articleRepo}

	// Initialize controllers
	authController := &controller.AuthController{Service: authService}
	articleController := &controller.ArticleController{Service: articleService}

	// Public routes (no authentication required)
	r.POST("/signup", authController.Signup)
	r.POST("/login", authController.Login)

	// Protected routes (authentication required)
	protected := r.Group("/")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.POST("/articles", articleController.CreateArticle)
		protected.GET("/articles", articleController.GetArticles)
	}

	log.Println("✓ Routes configured successfully")
	return nil
}