package user

import (
	"trabalho-01-batalha-naval/domain/entity"
	"trabalho-01-batalha-naval/domain/service"
)

type SignInUseCase struct {
	userService *service.UserService
}

func NewSignInUseCase(userService *service.UserService) *SignInUseCase {
	return &SignInUseCase{
		userService,
	}
}

func (u *SignInUseCase) Execute(username string, password string) (*entity.User, error) {
	return u.userService.SignIn(username, password)
}
