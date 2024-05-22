package handler

import (
	"bytesbanana/realworld-go-echo/src/internal/adapter/errs"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/kinbiko/jsonassert"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestUserHandler(t *testing.T) {

	assert := assert.New(t)
	ja := jsonassert.New(t)

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
				"token": "<<PRESENCE>>",
				"bio": null,
				"image": null
			}
		}`

		if assert.NoError(h.CreateUser(c)) {
			assert.Equal(http.StatusCreated, rec.Code)
			ja.Assertf(rec.Body.String(), expectedResponse)
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

	t.Run("given user information should return 422", func(t *testing.T) {

		rec, c := setup(func() *http.Request {
			userJSON := `{
			"user": {
				"email": "jake@test.com",
				"username": "jake",
				"password": "password"
			}
		}`
			req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			return req
		})

		userService := &StubUserSerivce{err: errs.ErrAlreadyBeenTaken}
		h := New(userService)

		if assert.NoError(h.CreateUser(c)) {
			expectedResponse := `{
				"errors": { "email": [ "has already been taken"]}
			  }`
			assert.Equal(http.StatusUnprocessableEntity, rec.Code)
			assert.JSONEq(expectedResponse, rec.Body.String())
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
