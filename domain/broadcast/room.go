package broadcast

import "trabalho-01-batalha-naval/domain/entity"

type RoomBroadcast interface {
	NotifyRoomCreated(room *entity.Room, createdBy *entity.User)

	NotifyUserIngressed(room *entity.Room, createdBy *entity.User)

	NotifyUserEgressed(room *entity.Room, user *entity.User)

	NotifyUserReady(room *entity.Room, user *entity.User)

	NotifyRoomDeleted(room *entity.Room)

	NotifyRoomStarted(room *entity.Room)

	NotifyUserFired(room *entity.Room, roomMove *entity.RoomMove)

	NotifyRoomCurrentPlayerChanged(room *entity.Room, user *entity.User)

	NotifyUserWon(room *entity.Room, user *entity.User)
}
