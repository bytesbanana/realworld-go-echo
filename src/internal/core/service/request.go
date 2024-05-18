package service

type UserCreateRequest struct {
	User struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	} `json:"user"`
}
