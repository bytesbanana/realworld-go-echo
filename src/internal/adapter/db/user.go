package db

import (
	"bytesbanana/realworld-go-echo/src/internal/core/domain"

	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

type userRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *userRepository {

	return &userRepository{
		db: db,
	}

}

func (ur *userRepository) CreateUser(email, username, password string) (*domain.User, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	rows, err := ur.db.NamedQuery("INSERT INTO users (email, username, hashed_password) VALUES (:email, :username, :password) RETURNING *",
		map[string]interface{}{
			"email":    email,
			"username": username,
			"password": hashedPassword,
		})
	if err != nil {
		return nil, err
	}

	u := new(domain.User)

	if rows.Next() {
		if err := rows.StructScan(&u); err != nil {
			return nil, err
		}
	}

	return u, nil
}
