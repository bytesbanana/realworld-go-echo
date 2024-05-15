package handler

import "bytesbanana/realworld-go-echo/internal/core/port"

type Handler struct {
	us port.UserService
}

func New(us port.UserService) Handler {
	return Handler{
		us: us,
	}
}
