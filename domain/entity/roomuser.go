package entity

import (
	"time"
)

type RoomUser struct {
	RoomUUID    string
	UserUUID    string
	ShipsSchema *Schema
	JoinedAt    time.Time
	WonAt       *time.Time
	LostAt      *time.Time
	AbandonedAt *time.Time
}

func (ru *RoomUser) HasSchema() bool {
	return ru.ShipsSchema != nil
}

func (ru *RoomUser) IsReady() bool {
	return ru.HasSchema()
}
