package service

import (
	"BlogPost/models"
	"BlogPost/repository"
	"BlogPost/utils"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	UserRepo *repository.UserRepository
}

func (s *AuthService) Signup(user models.User) error {
	hash, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	user.Password = string(hash)

	return s.UserRepo.CreateUser(user)
}

func (s *AuthService) Login(email, password string) (string, error) {
	user, err := s.UserRepo.GetUserByEmail(email)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", err
	}

	return utils.GenerateToken(user.UserID)
}
