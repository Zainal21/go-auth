package repository

import (
	"database/sql"
	"go-auth/entity"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db}
}

func (r UserRepository) FindUserByUserName(username string) (*entity.User, error) {
	query := "SELECT id, username, password FROM users WHERE username = ?"
	row := r.db.QueryRow(query, username)

	var user entity.User
	err := row.Scan(&user.ID, &user.Username, &user.Password)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) CreateUser(user *entity.User) error {
	query := "INSERT INTO users (username, password) VALUES (?, ?)"
	_, err := r.db.Exec(query, user.Username, user.Password)
	if err != nil {
		return err
	}
	return nil
}
