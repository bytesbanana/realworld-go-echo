package service

import "bytesbanana/realworld-go-echo/src/internal/core/domain"

type StubUserRepository struct {
	err error
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

	return &domain.User{
		ID:             1,
		Email:          email,
		Username:       "testuser",
		HashedPassword: "password",
	}, nil
}
