package service

type UserService interface {
	Register(req *UserCreateRequest) (*UserResponse, error)
}
