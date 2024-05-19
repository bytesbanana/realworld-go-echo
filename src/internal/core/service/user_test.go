package service

import (
	"bytesbanana/realworld-go-echo/src/internal/adapter/db"
	"bytesbanana/realworld-go-echo/src/internal/core/domain"
	"errors"

	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

type StubUserRepository struct {
	err error
}

func (s *StubUserRepository) CreateUser(email, username, password string) (*domain.User, error) {
	if s.err != nil {
		return nil, s.err
	}

	return &domain.User{
		ID:             1,
		Email:          email,
		Username:       username,
		HashedPassword: password,
	}, nil
}

func TestUserSerive(t *testing.T) {

	assert := assert.New(t)

	t.Run("should return user service instance", func(t *testing.T) {
		mockDB, _, err := sqlmock.New()
		assert.NoError(err)

		dbx := sqlx.NewDb(mockDB, "sqlmock")
		ur := db.NewUserRepository(dbx)
		s := NewUserService(ur)

		assert.NotNil(s)
		assert.Equal(ur, s.ur)
	})

	t.Run("should register user", func(t *testing.T) {

		s := NewUserService(&StubUserRepository{})
		assert.NotNil(s)

		res, err := s.Register(UserCreateRequest{
			User: struct {
				Username string `json:"username"`
				Email    string `json:"email"`
				Password string `json:"password"`
			}{
				Username: "testuser",
				Email:    "testuser@test.com",
				Password: "password",
			},
		})

		assert.NoError(err)
		assert.NotNil(res)
		assert.Equal("testuser", res.User.Username)
		assert.Equal("testuser@test.com", res.User.Email)
		assert.Equal("", res.User.Bio)
		assert.Equal("", res.User.Image)

	})

	t.Run("should return error when unable to create user", func(t *testing.T) {
		s := NewUserService(&StubUserRepository{
			err: errors.New("unable to create user"),
		})

		assert.NotNil(s)

		res, err := s.Register(UserCreateRequest{
			User: struct {
				Username string `json:"username"`
				Email    string `json:"email"`
				Password string `json:"password"`
			}{
				Username: "testuser",
				Email:    "testuser@test.com",
				Password: "password",
			},
		})

		assert.Error(err)
		assert.Nil(res)

	})

}
