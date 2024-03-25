package controller

import (
	"trabalho-01-batalha-naval/domain/service"
)

type RoomController struct {
	service *service.RoomService
}

func NewRoomController(service *service.RoomService) *RoomController {
	return &RoomController{
		service,
	}
}
