package repository

import (
	"database/sql"
	"go-auth/entity"
	"go-auth/utils"

	"golang.org/x/crypto/bcrypt"
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
	hashPassword, err := utils.HashPassword(user.Password)

	if err != nil {
		return err
	}

	query := "INSERT INTO users (username, password) VALUES (?, ?)"
	_, err = r.db.Exec(query, user.Username, hashPassword)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) VerifyPassword(username, password string) bool {
	storedHashedPassword, err := r.getHashedPassword(username)
	if err != nil {
		return false
	}

	err = bcrypt.CompareHashAndPassword([]byte(storedHashedPassword), []byte(password))
	if err != nil {
		return false
	}

	return true
}

func (r *UserRepository) getHashedPassword(username string) (string, error) {
	query := "SELECT password FROM users WHERE username = ?"
	row := r.db.QueryRow(query, username)

	var storedHashedPassword string
	err := row.Scan(&storedHashedPassword)
	if err != nil {
		return "", err
	}

	return storedHashedPassword, nil
}
