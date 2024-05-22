package service

import (
	"bytesbanana/realworld-go-echo/src/internal/core/domain"
	"bytesbanana/realworld-go-echo/src/internal/utils"
)

type UserResponse struct {
	User struct {
		Username string  `json:"username"`
		Email    string  `json:"email"`
		Token    string  `json:"token"`
		Bio      *string `json:"bio"`
		Image    *string `json:"image"`
	} `json:"user"`
}

func NewUserResponse(u *domain.User) (*UserResponse, error) {
	res := new(UserResponse)
	res.User.Username = u.Username
	res.User.Email = u.Email

	token, err := utils.GenerateToken(u.Username, u.Email)
	if err != nil {
		return nil, err
	}

	res.User.Token = token

	return res, nil
}
