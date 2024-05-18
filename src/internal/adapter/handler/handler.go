package handler

import "bytesbanana/realworld-go-echo/src/internal/core/service"

type Handler struct {
	us service.UserService
}

func New(us service.UserService) Handler {
	return Handler{
		us: us,
	}
}
