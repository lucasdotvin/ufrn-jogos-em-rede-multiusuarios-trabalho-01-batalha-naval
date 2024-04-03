package user

import (
	"trabalho-01-batalha-naval/domain/entity"
	"trabalho-01-batalha-naval/domain/service"
)

type SignUpUseCase struct {
	userService *service.UserService
}

func NewSignUpUseCase(userService *service.UserService) *SignUpUseCase {
	return &SignUpUseCase{
		userService,
	}
}

func (u *SignUpUseCase) Execute(name string, username string, password string) (*entity.User, error) {
	return u.userService.SignUp(name, username, password)
}
