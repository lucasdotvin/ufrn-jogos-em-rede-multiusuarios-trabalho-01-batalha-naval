package entity

import (
	"errors"
	"time"
)

type Room struct {
	UUID                 string
	Name                 string
	MapHeight            int
	MapWidth             int
	MaxPlayers           int
	CurrentPlayers       int
	ReadyPlayers         int
	UserCurrentlyPlaying *string
	CreatedBy            string
	CreatedAt            time.Time
	UpdatedAt            *time.Time
	StartedAt            *time.Time
	FinishedAt           *time.Time
}

func (r *Room) CanIngress() bool {
	return r.CurrentPlayers < r.MaxPlayers
}

func (r *Room) IsStarted() bool {
	return r.StartedAt != nil
}

func (r *Room) IsFinished() bool {
	return r.FinishedAt != nil
}

func (r *Room) IsEmpty() bool {
	return r.CurrentPlayers == 0
}

func (r *Room) IsActive() bool {
	return !r.IsFinished()
}

func (r *Room) IsFull() bool {
	return r.CurrentPlayers == r.MaxPlayers
}

func (r *Room) CanStart() bool {
	return r.CurrentPlayers == r.MaxPlayers && r.ReadyPlayers == r.MaxPlayers
}

var (
	RoomIsFullError                  = errors.New("room is full")
	RoomNotFoundError                = errors.New("room not found")
	RoomNotStartedError              = errors.New("room not started")
	UserNotPlayingError              = errors.New("user not playing")
	UserAlreadyInRoomError           = errors.New("user already in a room")
	UserNotInRoomError               = errors.New("user not in a room")
	UserAlreadyRegisteredSchemaError = errors.New("user already registered schema")
	InvalidSchemaError               = errors.New("invalid schema")
	RoomHasNoneActiveUserError       = errors.New("room has none active user")
	InvalidFirePositionError         = errors.New("invalid fire position")
)
