package controller

import (
	"BlogPost/models"
	"BlogPost/service"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	Service *service.AuthService
}

func (c *AuthController) Signup(ctx *gin.Context) {
	var user models.User
	ctx.BindJSON(&user)

	err := c.Service.Signup(user)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": "User created"})
}

func (c *AuthController) Login(ctx *gin.Context) {
	var input struct {
		Email    string
		Password string
	}

	ctx.BindJSON(&input)

	token, err := c.Service.Login(input.Email, input.Password)
	if err != nil {
		ctx.JSON(401, gin.H{"error": "Invalid credentials"})
		return
	}

	ctx.JSON(200, gin.H{"token": token})
}
