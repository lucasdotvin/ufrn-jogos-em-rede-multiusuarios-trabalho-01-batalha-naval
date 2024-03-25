package sqlite

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"trabalho-01-batalha-naval/infrastructure/config"
)

const driver = "sqlite3"

var cachedDb *sql.DB

func NewDatabase(cfg config.Config) (*sql.DB, error) {
	if cachedDb != nil {
		return cachedDb, nil
	}

	file, err := os.OpenFile(cfg.SQLiteDatabasePath, os.O_CREATE|os.O_RDWR, 0644)

	if err != nil {
		return nil, err
	}

	err = file.Close()

	if err != nil {
		return nil, err
	}

	db, err := sql.Open(driver, cfg.SQLiteDatabasePath)

	if err != nil {
		return nil, err
	}

	err = runMigrations(db)

	if err != nil {
		return nil, err
	}

	cachedDb = db

	return db, nil
}

func runMigrations(db *sql.DB) error {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS rooms (
			uuid TEXT PRIMARY KEY NOT NULL,
			name TEXT NOT NULL,
			player_1_uuid TEXT DEFAULT NULL,
			player_1_placements TEXT DEFAULT NULL,
			player_2_uuid TEXT DEFAULT NULL,
			player_2_placements TEXT DEFAULT NULL,
			last_move_player TEXT DEFAULT NULL,
			last_move_at DATETIME DEFAULT NULL,
			created_at DATETIME NOT NULL,
			updated_at DATETIME DEFAULT NULL,
			started_at DATETIME DEFAULT NULL
		);
	`)

	if err != nil {
		return err
	}

	return nil
}
