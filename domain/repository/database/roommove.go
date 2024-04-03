package database

import "trabalho-01-batalha-naval/domain/entity"

type RoomMoveRepository interface {
	Store(roomUser *entity.RoomMove) error

	CountHits(roomUUID string, userUUID string) (int, error)
}
