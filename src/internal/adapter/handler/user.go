package handler

import (
	"bytesbanana/realworld-go-echo/src/internal/core/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h Handler) CreateUser(c echo.Context) error {

	req := new(service.UserCreateRequest)

	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	u, err := h.us.Register(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, u)
}
