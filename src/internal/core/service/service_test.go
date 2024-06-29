package service

import (
	"bytesbanana/realworld-go-echo/src/internal/core/domain"
	"errors"
)

type StubUserRepository struct {
	users []domain.User
	err   error
}

func (s *StubUserRepository) CreateUser(email, username, password string) (*domain.User, error) {
	if s.err != nil {
		return nil, s.err
	}

	return &domain.User{
		ID:             1,
		Email:          email,
		Username:       username,
		HashedPassword: password,
	}, nil
}

func (s *StubUserRepository) GetUserByEmail(email string) (*domain.User, error) {
	if s.err != nil {
		return nil, s.err
	}

	for _, u := range s.users {
		if u.Email == email {
			return &u, nil
		}
	}

	return nil, errors.New("User not found")
}
