package user

import (
	"trabalho-01-batalha-naval/domain/entity"
	"trabalho-01-batalha-naval/domain/service"
)

type FindUserUseCase struct {
	userService *service.UserService
}

func NewFindUserUseCase(userService *service.UserService) *FindUserUseCase {
	return &FindUserUseCase{
		userService,
	}
}

func (u *FindUserUseCase) Execute(uuid string) (*entity.User, error) {
	return u.userService.Find(uuid)
}
