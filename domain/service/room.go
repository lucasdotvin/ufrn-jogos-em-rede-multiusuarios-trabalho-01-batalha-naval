package service

import "trabalho-01-batalha-naval/domain/repository"

type RoomService struct {
	repository repository.RoomRepository
}

func NewRoomService(repository repository.RoomRepository) *RoomService {
	return &RoomService{
		repository,
	}
}
