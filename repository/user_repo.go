package repository

import (
	"BlogPost/models"
	"database/sql"
)

type UserRepository struct {
	DB *sql.DB
}

func (r *UserRepository) CreateUser(user models.User) error {
	_, err := r.DB.Exec(
		"INSERT INTO users (name, email, password) VALUES (?, ?, ?)",
		user.Name, user.Email, user.Password,
	)
	return err
}

func (r *UserRepository) GetUserByEmail(email string) (models.User, error) {
	var user models.User

	err := r.DB.QueryRow(
		"SELECT user_id, name, email, password FROM users WHERE email = ?",
		email,
	).Scan(&user.UserID, &user.Name, &user.Email, &user.Password)

	return user, err
}
