package room

import (
	"trabalho-01-batalha-naval/config"
	"trabalho-01-batalha-naval/domain/entity"
	"trabalho-01-batalha-naval/domain/service"
)

type CreateRoomUseCase struct {
	cfg         config.Config
	roomService *service.RoomService
}

func NewCreateRoomUseCase(cfg config.Config, roomService *service.RoomService) *CreateRoomUseCase {
	return &CreateRoomUseCase{
		cfg,
		roomService,
	}
}

func (u *CreateRoomUseCase) Execute(name string, creatorUuid string) (*entity.Room, *entity.User, error) {
	ro, us, err := u.roomService.Create(name, u.cfg.MapHeight, u.cfg.MapWidth, u.cfg.DefaultMaxPlayers, creatorUuid)

	if err != nil {
		return nil, nil, err
	}

	return ro, us, nil
}
