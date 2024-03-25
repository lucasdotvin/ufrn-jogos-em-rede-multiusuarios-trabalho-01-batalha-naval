package repository

import "trabalho-01-batalha-naval/domain/entity"

type RoomRepository interface {
	Create(name string) (*entity.Room, error)
}
