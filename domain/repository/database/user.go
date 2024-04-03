package database

import "trabalho-01-batalha-naval/domain/entity"

type UserRepository interface {
	CheckIfUsernameExists(username string) (bool, error)

	FindByUsername(username string) (*entity.User, error)

	FindByUuid(uuid string) (*entity.User, error)

	Store(user *entity.User) error

	GetWhereUuidIn(uuids []string) ([]*entity.User, error)
}
