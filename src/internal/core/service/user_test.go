package service

import (
	"bytesbanana/realworld-go-echo/src/internal/adapter/db"
	"bytesbanana/realworld-go-echo/src/internal/core/domain"
	"errors"

	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestUserService(t *testing.T) {
	t.Parallel()

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

		res, err := s.Register(&UserCreateRequest{
			User: struct {
				Username string `json:"username" validate:"required"`
				Email    string `json:"email" validate:"required,email"`
				Password string `json:"password" validate:"required"`
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
		assert.NotNil(res.User.Token, "token should not be null")
		assert.NotEmpty(res.User.Token, "token should not be empty")
		assert.Nil(res.User.Bio)
		assert.Nil(res.User.Image)

	})

	t.Run("should return error when unable to create user", func(t *testing.T) {
		s := NewUserService(&StubUserRepository{
			err: errors.New("unable to create user"),
		})

		assert.NotNil(s)

		res, err := s.Register(&UserCreateRequest{
			User: struct {
				Username string `json:"username" validate:"required"`
				Email    string `json:"email" validate:"required,email"`
				Password string `json:"password" validate:"required"`
			}{
				Username: "testuser",
				Email:    "testuser@test.com",
				Password: "password",
			},
		})

		assert.Error(err)
		assert.Nil(res)

	})

	t.Run("should login user", func(t *testing.T) {
		hashPwd, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
		s := NewUserService(&StubUserRepository{
			users: []domain.User{{
				Username:       "testuser",
				Email:          "testuser@test.com",
				HashedPassword: string(hashPwd),
			}},
		})
		assert.NotNil(s)

		res, err := s.Login(&UserLoginRequest{
			User: struct {
				Email    string `json:"email" validate:"required,email"`
				Password string `json:"password" validate:"required"`
			}{
				Email:    "testuser@test.com",
				Password: "password",
			},
		})

		assert.NoError(err)
		assert.NotNil(res)
		assert.Equal("testuser", res.User.Username)
		assert.Equal("testuser@test.com", res.User.Email)
		assert.NotNil(res.User.Token, "token should not be null")
		assert.NotEmpty(res.User.Token, "token should not be empty")
		assert.Nil(res.User.Bio)
		assert.Nil(res.User.Image)
	})

}
