package handler

import (
	"bytesbanana/realworld-go-echo/src/internal/core/service"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

type StubUserSerivce struct {
	err error
}

func (s *StubUserSerivce) Register(req *service.UserCreateRequest) (*service.UserResponse, error) {
	if s.err != nil {
		return nil, s.err
	}

	return &service.UserResponse{
		User: struct {
			Username string  `json:"username"`
			Email    string  `json:"email"`
			Bio      *string `json:"bio"`
			Image    *string `json:"image"`
		}{
			Username: req.User.Username,
			Email:    req.User.Email,
			Bio:      nil,
			Image:    nil,
		},
	}, nil
}

func TestUserHandler(t *testing.T) {

	assert := assert.New(t)

	t.Run("given user information should return 201", func(t *testing.T) {

		rec, c := setup(func() *http.Request {
			userJSON := `{
				"user":{
					"email": "jake@jake.jake",
					"username": "jake",
					"password": "password"
				}
			}`

			req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			return req
		})

		userService := &StubUserSerivce{}
		h := New(userService)

		expectedResponse := `{
			"user": {
				"email": "jake@jake.jake",
				"username": "jake",
				"bio": null,
				"image": null
			}
		}`

		if assert.NoError(h.CreateUser(c)) {
			assert.Equal(http.StatusCreated, rec.Code)
			assert.JSONEq(expectedResponse, rec.Body.String())
		}
	})

	t.Run("given invalid create user request should return 400", func(t *testing.T) {

		rec, c := setup(func() *http.Request {
			userJSON := `{
			"user": {
				"email": "jake",
				"name": "jake",
				"password": "password"
			}
		}`
			req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			return req
		})

		userService := &StubUserSerivce{}
		h := New(userService)

		if assert.NoError(h.CreateUser(c)) {
			assert.Equal(http.StatusBadRequest, rec.Code)
		}
	})

	t.Run("given empty request should return 400", func(t *testing.T) {

		rec, c := setup(func() *http.Request {
			req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader("{}"))
			return req
		})

		userService := &StubUserSerivce{}
		h := New(userService)

		if assert.NoError(h.CreateUser(c)) {
			assert.Equal(http.StatusBadRequest, rec.Code)
		}
	})

	t.Run("given user information should return 500 when database error", func(t *testing.T) {
		rec, c := setup(func() *http.Request {
			userJSON := `{
			"user":{
				"email": "jake@jake.jake",
				"username": "jake",
				"password": "password"
				}
			}`

			req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			return req
		})

		userService := &StubUserSerivce{
			err: errors.New("database error"),
		}
		h := New(userService)

		if assert.NoError(h.CreateUser(c)) {
			assert.Equal(http.StatusInternalServerError, rec.Code)
		}
	})
}
