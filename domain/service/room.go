package service

import (
	"fmt"
	"time"
	"trabalho-01-batalha-naval/domain/broadcast"
	"trabalho-01-batalha-naval/domain/entity"
	"trabalho-01-batalha-naval/domain/repository/database"
)

type RoomService struct {
	roomBroadcast broadcast.RoomBroadcast

	roomDatabaseRepository     database.RoomRepository
	roomMoveDatabaseRepository database.RoomMoveRepository
	roomUserDatabaseRepository database.RoomUserRepository
	userDatabaseRepository     database.UserRepository
}

func NewRoomService(
	roomBroadcast broadcast.RoomBroadcast,
	roomDatabaseRepository database.RoomRepository,
	roomMoveDatabaseRepository database.RoomMoveRepository,
	roomUserDataRepository database.RoomUserRepository,
	userDatabaseRepository database.UserRepository,
) *RoomService {
	return &RoomService{
		roomBroadcast,
		roomDatabaseRepository,
		roomMoveDatabaseRepository,
		roomUserDataRepository,
		userDatabaseRepository,
	}
}

func (s *RoomService) List() ([]*entity.Room, error) {
	return s.roomDatabaseRepository.GetAllOpen()
}

func (s *RoomService) Create(name string, mapHeight int, mapWidth int, maxPlayers int, createdBy string) (*entity.Room, *entity.User, error) {
	u, err := s.userDatabaseRepository.FindByUuid(createdBy)

	if err != nil {
		return nil, nil, err
	}

	if u == nil {
		return nil, nil, entity.UserNotFoundError
	}

	alreadyActiveRoom, err := s.roomUserDatabaseRepository.FindActiveRoomForUser(createdBy)

	if err != nil {
		return nil, nil, err
	}

	if alreadyActiveRoom != nil {
		return nil, nil, entity.UserAlreadyInRoomError
	}

	r := &entity.Room{
		Name:           name,
		MapHeight:      mapHeight,
		MapWidth:       mapWidth,
		MaxPlayers:     maxPlayers,
		CurrentPlayers: 1,
		CreatedBy:      createdBy,
		CreatedAt:      time.Now(),
	}

	err = s.roomDatabaseRepository.Store(r)

	if err != nil {
		return nil, nil, err
	}

	ru := &entity.RoomUser{
		RoomUUID: r.UUID,
		UserUUID: r.CreatedBy,
	}

	err = s.roomUserDatabaseRepository.Store(ru)

	if err != nil {
		return nil, nil, err
	}

	s.roomBroadcast.NotifyRoomCreated(r, u)

	return r, u, nil
}

func (s *RoomService) IngressUser(roomUuid string, userUuid string) (*entity.Room, error) {
	ro, err := s.roomDatabaseRepository.FindByUuid(roomUuid)

	if err != nil {
		return nil, err
	}

	if ro == nil {
		return nil, entity.RoomNotFoundError
	}

	if !ro.CanIngress() {
		return nil, entity.RoomIsFullError
	}

	us, err := s.userDatabaseRepository.FindByUuid(userUuid)

	if err != nil {
		return nil, err
	}

	if us == nil {
		return nil, entity.UserNotFoundError
	}

	alreadyActiveRoom, err := s.roomUserDatabaseRepository.FindActiveRoomForUser(userUuid)

	if err != nil {
		return nil, err
	}

	if alreadyActiveRoom != nil {
		return nil, entity.UserAlreadyInRoomError
	}

	ru := &entity.RoomUser{
		RoomUUID: ro.UUID,
		UserUUID: us.UUID,
	}

	err = s.roomUserDatabaseRepository.Store(ru)

	if err != nil {
		return nil, err
	}

	ro.CurrentPlayers++

	err = s.roomDatabaseRepository.Update(ro)

	if err != nil {
		return nil, err
	}

	s.roomBroadcast.NotifyUserIngressed(ro, us)

	return ro, nil
}

func (s *RoomService) ValidateUserSubscription(roomUuid string, userUuid string) error {
	ro, err := s.roomDatabaseRepository.FindByUuid(roomUuid)

	if err != nil {
		return err
	}

	if ro == nil {
		return entity.RoomNotFoundError
	}

	us, err := s.userDatabaseRepository.FindByUuid(userUuid)

	if err != nil {
		return err
	}

	if us == nil {
		return entity.UserNotFoundError
	}

	ru, err := s.roomUserDatabaseRepository.FindActiveRoomForUser(userUuid)

	if err != nil {
		return err
	}

	if ru == nil {
		return entity.UserNotInRoomError
	}

	if ru.RoomUUID != ro.UUID {
		return entity.UserNotInRoomError
	}

	return nil
}

func (s *RoomService) HandleUserDisconnect(userUuid string) error {
	ru, err := s.roomUserDatabaseRepository.FindActiveRoomForUser(userUuid)

	if err != nil {
		return err
	}

	if ru == nil {
		return nil
	}

	ro, err := s.roomDatabaseRepository.FindByUuid(ru.RoomUUID)

	if err != nil {
		return err
	}

	if ro == nil {
		return nil
	}

	if !ro.IsStarted() {
		return s.handleUserDisconnectBeforeGameStart(ru, ro)
	}

	return nil
}

func (s *RoomService) handleUserDisconnectBeforeGameStart(ru *entity.RoomUser, ro *entity.Room) error {
	err := s.roomUserDatabaseRepository.Delete(ru)

	if err != nil {
		return err
	}

	ro.CurrentPlayers--

	if ro.IsEmpty() {
		return s.deleteRoom(ro)
	}

	err = s.roomDatabaseRepository.Update(ro)

	if err != nil {
		return err
	}

	us, err := s.userDatabaseRepository.FindByUuid(ru.UserUUID)

	if err != nil {
		return err
	}

	if us == nil {
		return entity.UserNotFoundError
	}

	s.roomBroadcast.NotifyUserEgressed(ro, us)

	return nil
}

func (s *RoomService) FindUserActiveRoom(userUuid string) (*entity.Room, error) {
	ru, err := s.roomUserDatabaseRepository.FindActiveRoomForUser(userUuid)

	if err != nil {
		return nil, err
	}

	if ru == nil {
		return nil, nil
	}

	ro, err := s.roomDatabaseRepository.FindByUuid(ru.RoomUUID)

	if err != nil {
		fmt.Println("error while finding room by uuid: ", err)
		return nil, err
	}

	if ro == nil {
		fmt.Println("room not found")
		return nil, entity.RoomNotFoundError
	}

	if !ro.IsActive() {
		fmt.Println("room is active")
		return nil, entity.RoomNotFoundError
	}

	return ro, nil
}

func (s *RoomService) deleteRoom(room *entity.Room) error {
	err := s.roomDatabaseRepository.Delete(room)

	if err != nil {
		return err
	}

	s.roomBroadcast.NotifyRoomDeleted(room)

	return nil
}

func (s *RoomService) GetRoomUsers(roomUuid string) ([]*entity.RoomUser, error) {
	rus, err := s.roomUserDatabaseRepository.GetByRoomUuid(roomUuid)

	if err != nil {
		return nil, err
	}

	return rus, nil
}

func (s *RoomService) RegisterShipPlacements(roomUuid string, userUuid string, dispositions []*entity.Disposition) error {
	ro, err := s.roomDatabaseRepository.FindByUuid(roomUuid)

	if err != nil {
		return err
	}

	if ro == nil {
		return entity.RoomNotFoundError
	}

	ru, err := s.roomUserDatabaseRepository.FindActiveRoomForUser(userUuid)

	if err != nil {
		return err
	}

	if ru == nil {
		return entity.UserNotInRoomError
	}

	if ru.RoomUUID != ro.UUID {
		return entity.UserNotInRoomError
	}

	if ru.HasSchema() {
		return entity.UserAlreadyRegisteredSchemaError
	}

	us, err := s.userDatabaseRepository.FindByUuid(ru.UserUUID)

	if err != nil {
		return err
	}

	if us == nil {
		return entity.UserNotFoundError
	}

	schema := entity.NewSchema(dispositions)

	if !schema.IsValid(ro.MapHeight, ro.MapWidth) {
		return entity.InvalidSchemaError
	}

	ru.ShipsSchema = schema

	err = s.roomUserDatabaseRepository.Update(ru)

	if err != nil {
		return err
	}

	ro.ReadyPlayers++

	if ro.CanStart() {
		startedAt := time.Now()
		ro.StartedAt = &startedAt
	}

	err = s.roomDatabaseRepository.Update(ro)

	if err != nil {
		return err
	}

	if !ro.IsStarted() {
		s.roomBroadcast.NotifyUserReady(ro, us)

		return nil
	}

	err = s.handleRoomStarted(ro)

	if err != nil {
		return err
	}

	return nil
}

func (s *RoomService) handleRoomStarted(room *entity.Room) error {
	firstRoomUserToPlay, err := s.roomUserDatabaseRepository.FindRandomActivePlayer(room.UUID)

	if err != nil {
		return err
	}

	if firstRoomUserToPlay == nil {
		return entity.RoomHasNoneActiveUserError
	}

	startedAt := time.Now()
	room.StartedAt = &startedAt
	room.UserCurrentlyPlaying = &firstRoomUserToPlay.UserUUID

	err = s.roomDatabaseRepository.Update(room)

	if err != nil {
		return err
	}

	s.roomBroadcast.NotifyRoomStarted(room)

	return nil
}

func (s *RoomService) RegisterFire(roomUuid string, userUuid string, position *entity.Position) (*entity.RoomMove, error) {
	ro, err := s.roomDatabaseRepository.FindByUuid(roomUuid)

	if err != nil {
		fmt.Println("error while finding room by uuid: ", err)
		return nil, err
	}

	if ro == nil {
		fmt.Println("room not found")
		return nil, entity.RoomNotFoundError
	}

	if !ro.IsStarted() {
		fmt.Println("room not started")
		return nil, entity.RoomNotStartedError
	}

	if *ro.UserCurrentlyPlaying != userUuid {
		fmt.Println("user not playing")
		return nil, entity.UserNotPlayingError
	}

	rus, err := s.roomUserDatabaseRepository.GetByRoomUuid(roomUuid)

	if err != nil {
		fmt.Println("error while finding room users by room uuid: ", err)
		return nil, err
	}

	var currentRoomUser *entity.RoomUser
	var enemyRoomUser *entity.RoomUser

	for _, ru := range rus {
		if ru.UserUUID == userUuid {
			currentRoomUser = ru
		} else {
			enemyRoomUser = ru
		}
	}

	if currentRoomUser == nil {
		fmt.Println("current room user not found")
		return nil, entity.UserNotInRoomError
	}

	if enemyRoomUser == nil {
		fmt.Println("enemy room user not found")
		return nil, entity.RoomHasNoneActiveUserError
	}

	hit := enemyRoomUser.ShipsSchema.Hits(position)

	rm := &entity.RoomMove{
		RoomUUID: ro.UUID,
		UserUUID: userUuid,
		X:        position.X,
		Y:        position.Y,
		Hit:      hit,
	}

	err = s.roomMoveDatabaseRepository.Store(rm)

	if err != nil {
		fmt.Println("error while storing room move: ", err)
		return nil, entity.InvalidFirePositionError
	}

	s.roomBroadcast.NotifyUserFired(ro, rm)

	if !hit {
		return s.handleMissedFire(ro, rm, enemyRoomUser)
	}

	totalHits, err := s.roomMoveDatabaseRepository.CountHits(roomUuid, userUuid)

	if err != nil {
		fmt.Println("error while counting hits: ", err)
		return nil, err
	}

	if totalHits == entity.TotalShipsSize {
		return s.handleUserWin(ro, rm, currentRoomUser, enemyRoomUser)
	}

	return rm, nil
}

func (s *RoomService) handleMissedFire(room *entity.Room, roomMove *entity.RoomMove, enemyRoomUser *entity.RoomUser) (*entity.RoomMove, error) {
	room.UserCurrentlyPlaying = &enemyRoomUser.UserUUID

	err := s.roomDatabaseRepository.Update(room)

	if err != nil {
		fmt.Println("error while updating room: ", err)
		return nil, err
	}

	enemyUser, err := s.userDatabaseRepository.FindByUuid(enemyRoomUser.UserUUID)

	if err != nil {
		fmt.Println("error while finding user by uuid: ", err)
		return nil, err
	}

	s.roomBroadcast.NotifyRoomCurrentPlayerChanged(room, enemyUser)

	return roomMove, nil
}

func (s *RoomService) handleUserWin(room *entity.Room, roomMove *entity.RoomMove, currentRoomUser *entity.RoomUser, enemyRoomUser *entity.RoomUser) (*entity.RoomMove, error) {
	finishedAt := time.Now()
	room.FinishedAt = &finishedAt
	room.UserCurrentlyPlaying = nil

	err := s.roomDatabaseRepository.Update(room)

	if err != nil {
		fmt.Println("error while updating room: ", err)
		return nil, err
	}

	currentRoomUser.WonAt = &finishedAt

	err = s.roomUserDatabaseRepository.Update(currentRoomUser)

	if err != nil {
		fmt.Println("error while updating room user: ", err)
		return nil, err
	}

	enemyRoomUser.LostAt = &finishedAt

	err = s.roomUserDatabaseRepository.Update(enemyRoomUser)

	if err != nil {
		fmt.Println("error while updating room user: ", err)
		return nil, err
	}

	currentUser, err := s.userDatabaseRepository.FindByUuid(currentRoomUser.UserUUID)

	if err != nil {
		fmt.Println("error while finding user by uuid: ", err)
		return nil, err
	}

	s.roomBroadcast.NotifyUserWon(room, currentUser)

	return roomMove, nil
}
