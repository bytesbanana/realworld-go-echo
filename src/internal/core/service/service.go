package service

type UserService interface {
	Register(req *UserCreateRequest) (*UserResponse, error)
	Login(req *UserLoginRequest) (*UserResponse, error)
}
