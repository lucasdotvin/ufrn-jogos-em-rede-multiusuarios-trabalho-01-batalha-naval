package room

import (
	"trabalho-01-batalha-naval/domain/entity"
	"trabalho-01-batalha-naval/domain/service"
)

type RegisterShipPlacementsUseCase struct {
	roomService *service.RoomService
}

func NewRegisterShipPlacementsUseCase(roomService *service.RoomService) *RegisterShipPlacementsUseCase {
	return &RegisterShipPlacementsUseCase{
		roomService,
	}
}

func (u *RegisterShipPlacementsUseCase) Execute(roomUuid string, userUuid string, dispositions []*entity.Disposition) error {
	return u.roomService.RegisterShipPlacements(roomUuid, userUuid, dispositions)
}
