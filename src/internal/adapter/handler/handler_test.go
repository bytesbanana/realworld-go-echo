package handler

import (
	"net/http"
	"net/http/httptest"

	cv "bytesbanana/realworld-go-echo/src/internal/adapter/validator"
	"bytesbanana/realworld-go-echo/src/internal/core/service"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func setup(buildRequest func() *http.Request) (*httptest.ResponseRecorder, echo.Context) {
	e := echo.New()
	e.Validator = &cv.CustomValidator{Validator: validator.New()}
	req := buildRequest()
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	return rec, c

}

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
			Token    string  `json:"token"`
			Bio      *string `json:"bio"`
			Image    *string `json:"image"`
		}{
			Username: req.User.Username,
			Email:    req.User.Email,
			Token:    "jwt.token",
			Bio:      nil,
			Image:    nil,
		},
	}, nil
}
