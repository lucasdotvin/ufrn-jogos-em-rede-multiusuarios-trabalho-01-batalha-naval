package entity

import "time"

type RoomMove struct {
	RoomUUID  string
	UserUUID  string
	X         int
	Y         int
	Hit       bool
	CreatedAt time.Time
}
