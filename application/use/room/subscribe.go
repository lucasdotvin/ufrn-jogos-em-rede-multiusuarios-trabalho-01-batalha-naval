package room

import (
	"trabalho-01-batalha-naval/domain/service"
)

type ValidateUserSubscriptionUseCase struct {
	roomService *service.RoomService
}

func NewValidateUserSubscriptionUseCase(roomService *service.RoomService) *ValidateUserSubscriptionUseCase {
	return &ValidateUserSubscriptionUseCase{
		roomService,
	}
}

func (u *ValidateUserSubscriptionUseCase) Execute(roomUuid string, userUuid string) error {
	return u.roomService.ValidateUserSubscription(roomUuid, userUuid)
}
