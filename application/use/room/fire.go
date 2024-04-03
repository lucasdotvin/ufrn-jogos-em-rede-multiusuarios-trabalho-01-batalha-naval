package room

import (
	"trabalho-01-batalha-naval/domain/entity"
	"trabalho-01-batalha-naval/domain/service"
)

type RegisterFireUseCase struct {
	roomService *service.RoomService
}

func NewRegisterFireUseCase(roomService *service.RoomService) *RegisterFireUseCase {
	return &RegisterFireUseCase{
		roomService,
	}
}

func (u *RegisterFireUseCase) Execute(roomUuid string, userUuid string, position *entity.Position) (*entity.RoomMove, error) {
	return u.roomService.RegisterFire(roomUuid, userUuid, position)
}
