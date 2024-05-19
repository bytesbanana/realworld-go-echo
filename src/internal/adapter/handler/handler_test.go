package handler

import (
	"net/http"
	"net/http/httptest"

	cv "bytesbanana/realworld-go-echo/src/internal/adapter/validator"

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
