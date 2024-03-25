package main

import (
	"trabalho-01-batalha-naval/infrastructure/config"
	"trabalho-01-batalha-naval/infrastructure/database/driver/sqlite"
)

func main() {
	cfg := config.GetConfig()

	sqliteDb, err := sqlite.NewDatabase(cfg)

	if err != nil {
		panic("failed to connect database " + err.Error())
	}

	err = sqliteDb.Close()

	if err != nil {
		panic("failed to close database " + err.Error())
	}
}
