package db

import (
	"bytesbanana/realworld-go-echo/src/internal/adapter/errs"
	"database/sql"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestCreateUser(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)
	t.Run("given user should create user", func(t *testing.T) {
		mockDB, mock, err := sqlmock.New()
		assert.NoError(err, "an error '%s' was not expected when opening a stub database connection", err)

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
		assert.NoError(err)
		rows := sqlmock.NewRows([]string{"id", "email", "username", "hashed_password"}).
			AddRow(1, "testuser@test.com", "username", hashedPassword)

		mock.ExpectQuery("SELECT (.+) FROM users").WillReturnRows(sqlmock.NewRows([]string{}))
		mock.ExpectQuery("INSERT INTO users").WillReturnRows(rows)

		sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
		ur := NewUserRepository(sqlxDB)

		u, err := ur.CreateUser("testuser@test.com", "username", "password")
		assert.NoError(err)
		assert.NotNil(u, "user should not be null")
		assert.Equal(u.ID, 1)
		assert.Equal(u.Email, "testuser@test.com")
		assert.Equal(u.Username, "username")
		assert.NotEmpty(u.HashedPassword)
		assert.NotEqual("password", u.HashedPassword)
	})

	t.Run("given user should error when unable to insert to data", func(t *testing.T) {
		mockDB, mock, err := sqlmock.New()
		assert.NoError(err, "an error '%s' was not expected when opening a stub database connection", err)

		mock.ExpectQuery("SELECT (.+) FROM users").WillReturnError(sql.ErrNoRows)
		mock.ExpectQuery("INSERT INTO users").WithArgs("testuser@test.com", "username", "password").WillReturnError(errors.New("unable to insert user"))

		sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
		ur := NewUserRepository(sqlxDB)

		u, err := ur.CreateUser("testuser@test.com", "username", "password")
		assert.Error(err)
		assert.Nil(u)
	})

	t.Run("given user should error when unable to map struct", func(t *testing.T) {
		mockDB, mock, err := sqlmock.New()
		assert.NoError(err, "an error '%s' was not expected when opening a stub database connection", err)
		rows := sqlmock.NewRows([]string{"id", "email2", "username2", "hashed_password"}).
			AddRow(1, "testuser@test.com", "username", "password")

		mock.ExpectQuery("SELECT (.+) FROM users").WillReturnError(sql.ErrNoRows)
		mock.ExpectQuery("INSERT INTO users").WithArgs("testuser@test.com", "username", "password").WillReturnRows(rows)

		sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
		ur := NewUserRepository(sqlxDB)

		u, err := ur.CreateUser("testuser@test.com", "username", "password")
		assert.Error(err)
		assert.Nil(u)
	})

	t.Run("giver user information should return error when user already exists", func(t *testing.T) {
		mockDB, mock, err := sqlmock.New()
		assert.NoError(err, "an error '%s' was not expected when opening a stub database connection", err)

		columns := []string{"id", "email", "username", "hashed_password"}
		rows := sqlmock.NewRows(columns)
		rows.AddRow(1, "testuser@test.com", "username", "password")

		mock.ExpectQuery("SELECT (.+) FROM users").WillReturnRows(rows)

		sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
		ur := NewUserRepository(sqlxDB)

		u, err := ur.CreateUser("testuser@test.com", "username", "password")
		assert.Error(err)
		assert.Nil(u)
		assert.EqualError(err, errs.ErrAlreadyBeenTaken.Error())
	})
}

func TestGetUserByEmail(t *testing.T) {
	assert := assert.New(t)

	t.Run("given existing email should return user", func(t *testing.T) {
		mockDB, mock, err := sqlmock.New()
		assert.NoError(err, "an error '%s' was not expected when opening a stub database connection", err)
		rows := sqlmock.NewRows([]string{"id", "email", "username", "hashed_password", "bio", "image"}).
			AddRow(1, "testuser@test.com", "username", "hashed_password", nil, nil)

		mock.ExpectQuery("SELECT (.+) FROM users").WithArgs("testuser@test.com").WillReturnRows(rows)

		sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
		ur := NewUserRepository(sqlxDB)

		u, err := ur.GetUserByEmail("testuser@test.com")
		assert.NoError(err)
		assert.NotNil(u)
		assert.Equal(1, u.ID)
		assert.Equal("testuser@test.com", u.Email)
		assert.Equal("username", u.Username)
		assert.Equal("hashed_password", u.HashedPassword)
		assert.Nil(u.Bio)
		assert.Nil(u.Image)

	})
}
