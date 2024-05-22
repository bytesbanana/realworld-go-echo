package service

import (
	"bytesbanana/realworld-go-echo/src/internal/core/port"
)

type userService struct {
	ur port.UserRepository
}

func NewUserService(ur port.UserRepository) *userService {
	return &userService{
		ur: ur,
	}
}

func (s *userService) Register(req *UserCreateRequest) (*UserResponse, error) {

	u, err := s.ur.CreateUser(req.User.Email, req.User.Username, req.User.Password)
	if err != nil {
		return nil, err
	}

	return NewUserResponse(u), nil
}
