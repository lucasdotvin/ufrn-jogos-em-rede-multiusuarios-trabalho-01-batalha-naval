package sqlite

import (
	"database/sql"
	"time"
	"trabalho-01-batalha-naval/domain/entity"
)

type RoomMoveRepository struct {
	db *sql.DB
}

func NewSqliteRoomMoveRepository(db *sql.DB) *RoomMoveRepository {
	return &RoomMoveRepository{
		db,
	}
}

func (s *RoomMoveRepository) Store(roomMove *entity.RoomMove) error {
	createdAt := time.Now()

	_, err := s.db.Exec(`
		INSERT INTO room_moves (room_uuid, user_uuid, x, y, hit, created_at)
		VALUES (?, ?, ?, ?, ?, ?)
	`, roomMove.RoomUUID, roomMove.UserUUID, roomMove.X, roomMove.Y, roomMove.Hit, createdAt)

	if err != nil {
		return err
	}

	roomMove.CreatedAt = createdAt

	return nil
}

func (s *RoomMoveRepository) CountHits(roomUUID string, userUUID string) (int, error) {
	row := s.db.QueryRow(`
		SELECT COUNT(*) FROM room_moves
		WHERE
		    room_uuid = ?
		  	AND user_uuid = ?
		    AND hit = true
	`, roomUUID, userUUID)

	var count int
	err := row.Scan(&count)

	if err != nil {
		return 0, err
	}

	return count, nil
}
