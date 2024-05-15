package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h Handler) CreateUser(c echo.Context) error {

	var req UserCreateRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	u, err := h.us.Register(req.User.Email, req.User.Username, req.User.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, u)
}
