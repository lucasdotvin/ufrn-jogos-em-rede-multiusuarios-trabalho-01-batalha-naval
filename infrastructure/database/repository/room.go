package repository

import (
	"database/sql"
	"github.com/google/uuid"
	"time"
	"trabalho-01-batalha-naval/domain/entity"
)

type RoomRepository struct {
	db *sql.DB
}

func NewRoomRepository(db *sql.DB) *RoomRepository {
	return &RoomRepository{
		db,
	}
}

func (r *RoomRepository) Create(name string) (*entity.Room, error) {
	u, err := uuid.NewRandom()

	if err != nil {
		return nil, err
	}

	cratedAt := time.Now()

	_, err = r.db.Exec(`
		INSERT INTO rooms (uuid, name, created_at)
		VALUES (?, ?, ?)
	`, u.String(), name, cratedAt)

	if err != nil {
		return nil, err
	}

	return &entity.Room{
		Uuid:      u.String(),
		Name:      name,
		CreatedAt: cratedAt,
	}, nil
}

func (r *RoomRepository) GetAll() ([]*entity.Room, error) {
	rows, err := r.db.Query(`
		SELECT uuid, name, created_at
		FROM rooms
	`)

	if err != nil {
		return nil, err
	}

	var rooms []*entity.Room

	for rows.Next() {
		var room entity.Room

		err = rows.Scan(&room.Uuid, &room.Name, &room.CreatedAt)

		if err != nil {
			return nil, err
		}

		rooms = append(rooms, &room)
	}

	err = rows.Close()

	if err != nil {
		return nil, err
	}

	return rooms, nil
}
