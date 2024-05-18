package port

import (
	"bytesbanana/realworld-go-echo/src/internal/core/domain"
)

type UserRepository interface {
	CreateUser(email, username, password string) (*domain.User, error)
}
