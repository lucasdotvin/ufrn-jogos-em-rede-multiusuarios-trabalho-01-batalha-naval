package roommove

import (
	"time"
	"trabalho-01-batalha-naval/domain/entity"
)

type RoomMoveResponse struct {
	RoomUUID  string    `json:"room_uuid"`
	UserUUID  string    `json:"user_uuid"`
	X         int       `json:"x"`
	Y         int       `json:"y"`
	Hit       bool      `json:"hit"`
	CreatedAt time.Time `json:"created_at"`
}

func NewRoomMoveResponse(roomMove *entity.RoomMove) *RoomMoveResponse {
	return &RoomMoveResponse{
		RoomUUID:  roomMove.RoomUUID,
		UserUUID:  roomMove.UserUUID,
		X:         roomMove.X,
		Y:         roomMove.Y,
		Hit:       roomMove.Hit,
		CreatedAt: roomMove.CreatedAt,
	}
}
