package database

import "trabalho-01-batalha-naval/domain/entity"

type RoomRepository interface {
	GetAllOpen() ([]*entity.Room, error)

	Store(room *entity.Room) error

	Update(room *entity.Room) error

	FindByUuid(uuid string) (*entity.Room, error)

	Delete(room *entity.Room) error
}
