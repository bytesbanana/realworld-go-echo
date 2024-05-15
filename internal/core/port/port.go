package port

import (
	"bytesbanana/realworld-go-echo/internal/adapter/dto"
	"bytesbanana/realworld-go-echo/internal/core/domain"
)

type UserRepository interface {
	CreateUser(email, username, password string) (*domain.User, error)
}

type UserService interface {
	Register(email, username, password string) (*dto.UserResponse, error)
}
