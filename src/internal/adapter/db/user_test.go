package db

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {

	t.Run("given user should create user", func(t *testing.T) {
		mockDB, mock, err := sqlmock.New()
		assert := assert.New(t)
		assert.NoError(err, "an error '%s' was not expected when opening a stub database connection", err)

		rows := sqlmock.NewRows([]string{"id", "email", "username", "hashed_password"}).
			AddRow(1, "testuser@test.com", "username", "password")

		mock.ExpectQuery("INSERT INTO users").WithArgs("testuser@test.com", "username", "password").WillReturnRows(rows)

		sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
		ur := NewUserRepository(sqlxDB)

		u, err := ur.CreateUser("testuser@test.com", "username", "password")
		assert.NoError(err)
		assert.NotNil(u, "user should not be null")
		assert.Equal(u.ID, 1)
		assert.Equal(u.Email, "testuser@test.com")
		assert.Equal(u.Username, "username")
		assert.NotEmpty(u.HashedPassword)
	})

	t.Run("given user should error", func(t *testing.T) {
		mockDB, mock, err := sqlmock.New()
		assert := assert.New(t)
		assert.NoError(err, "an error '%s' was not expected when opening a stub database connection", err)

		mock.ExpectQuery("INSERT INTO users").WithArgs("testuser@test.com", "username", "password").WillReturnError(errors.New("unable to insert user"))

		sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
		ur := NewUserRepository(sqlxDB)

		u, err := ur.CreateUser("testuser@test.com", "username", "password")
		assert.Error(err)
		assert.Nil(u)
	})
}
