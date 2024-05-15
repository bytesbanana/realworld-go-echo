package service

import (
	"bytesbanana/realworld-go-echo/internal/adapter/dto"
	"bytesbanana/realworld-go-echo/internal/core/port"
)

type userService struct {
	ur port.UserRepository
}

func NewUserService(ur port.UserRepository) *userService {
	return &userService{
		ur: ur,
	}
}

func (s *userService) Register(email, username, password string) (*dto.UserResponse, error) {

	u, err := s.ur.CreateUser(email, username, password)
	if err != nil {
		return nil, err
	}

	return dto.NewUserResponse(u), nil
}
