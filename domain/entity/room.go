package entity

import "time"

type Room struct {
	Uuid              string
	Name              string
	Player1Uuid       *string
	Player1Placements []ShipPlacement
	Player2Uuid       *string
	Player2Placements []ShipPlacement
	LastMovePlayer    *Player
	LastMoveAt        *time.Time
	CreatedAt         time.Time
	UpdatedAt         *time.Time
	StartedAt         *time.Time
}
