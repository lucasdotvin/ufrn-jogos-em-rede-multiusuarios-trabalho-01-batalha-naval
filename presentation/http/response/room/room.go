package room

import (
	"trabalho-01-batalha-naval/domain/entity"
	"trabalho-01-batalha-naval/presentation/http/response"
	"trabalho-01-batalha-naval/presentation/http/response/user"
)

type RoomResponse struct {
	UUID                 string            `json:"uuid"`
	Name                 string            `json:"name"`
	MapHeight            int               `json:"map_height"`
	MapWidth             int               `json:"map_width"`
	MaxPlayers           int               `json:"max_players"`
	CurrentPlayers       int               `json:"current_players"`
	ReadyPlayers         int               `json:"ready_players"`
	UserCurrentlyPlaying *string           `json:"user_currently_playing"`
	CreatedBy            *user.UserField   `json:"created_by"`
	Users                []*user.UserField `json:"users"`
	CreatedAt            string            `json:"created_at"`
	UpdatedAt            *string           `json:"updated_at"`
	StartedAt            *string           `json:"started_at"`
	FinishedAt           *string           `json:"finished_at"`
}

func NewRoomResponse(room *entity.Room, createdBy *entity.User, users []*entity.User, roomUsers []*entity.RoomUser) *RoomResponse {
	if room == nil {
		return nil
	}

	return &RoomResponse{
		UUID:                 room.UUID,
		Name:                 room.Name,
		MapHeight:            room.MapHeight,
		MapWidth:             room.MapWidth,
		MaxPlayers:           room.MaxPlayers,
		CurrentPlayers:       room.CurrentPlayers,
		ReadyPlayers:         room.ReadyPlayers,
		UserCurrentlyPlaying: room.UserCurrentlyPlaying,
		CreatedBy:            user.NewUserField(createdBy, nil),
		Users:                user.NewUsersField(users, roomUsers),
		CreatedAt:            response.FormatTimeField(room.CreatedAt),
		UpdatedAt:            response.FormatOptionalTimeField(room.UpdatedAt),
		StartedAt:            response.FormatOptionalTimeField(room.StartedAt),
		FinishedAt:           response.FormatOptionalTimeField(room.FinishedAt),
	}
}
