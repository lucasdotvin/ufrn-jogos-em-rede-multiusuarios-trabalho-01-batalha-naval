package roomuser

import "trabalho-01-batalha-naval/domain/entity"

type RoomUserField struct {
	IsReady bool `json:"is_ready"`
}

func NewRoomUserField(roomUser *entity.RoomUser) *RoomUserField {
	if roomUser == nil {
		return nil
	}

	return &RoomUserField{
		IsReady: roomUser.IsReady(),
	}
}
