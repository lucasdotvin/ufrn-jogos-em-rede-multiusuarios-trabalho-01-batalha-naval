package database

import "trabalho-01-batalha-naval/domain/entity"

type RoomUserRepository interface {
	Store(roomUser *entity.RoomUser) error

	FindActiveRoomForUser(userUUID string) (*entity.RoomUser, error)

	Delete(roomUser *entity.RoomUser) error

	GetByRoomUuid(roomUUID string) ([]*entity.RoomUser, error)

	Update(roomUser *entity.RoomUser) error

	FindRandomActivePlayer(roomUUID string) (*entity.RoomUser, error)
}
