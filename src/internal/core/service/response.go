package service

import "bytesbanana/realworld-go-echo/src/internal/core/domain"

type UserResponse struct {
	User struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Bio      string `json:"bio"`
		Image    string `json:"image"`
	} `json:"user"`
}

func NewUserResponse(u *domain.User) *UserResponse {
	res := new(UserResponse)
	res.User.Username = u.Username
	res.User.Email = u.Email

	return res
}
