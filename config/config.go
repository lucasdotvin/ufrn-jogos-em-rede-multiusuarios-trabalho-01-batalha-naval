package config

import (
	"github.com/joho/godotenv"
	"os"
	"strconv"
)

type Config struct {
	DefaultMaxPlayers       int
	HashCost                int
	JwtDurationInMinutes    int
	JwtRenewDueInMinutes    int
	JwtSecret               string
	MapHeight               int
	MapWidth                int
	ServerAddress           string
	SQLiteDatabasePath      string
	WebSocketAllowedOrigins string
}

var config *Config = nil

func GetConfig() Config {
	return *config
}

func init() {
	err := godotenv.Load()

	if err != nil {
		panic("failed to load .env file")
	}

	parsedDefaultMaxPlayers, err := strconv.Atoi(os.Getenv("DEFAULT_MAX_PLAYERS"))

	if err != nil {
		panic("failed to parse DEFAULT_MAX_PLAYERS")
	}

	parsedHashCost, err := strconv.Atoi(os.Getenv("HASH_COST"))

	if err != nil {
		panic("failed to parse HASH_COST")
	}

	parsedJwtDurationInMinutes, err := strconv.Atoi(os.Getenv("JWT_DURATION_IN_MINUTES"))

	if err != nil {
		panic("failed to parse JWT_DURATION_IN_MINUTES")
	}

	parsedJwtRenewDueInMinutes, err := strconv.Atoi(os.Getenv("JWT_RENEW_DUE_IN_MINUTES"))

	if err != nil {
		panic("failed to parse JWT_RENEW_DUE_IN_MINUTES")
	}

	parsedMapHeight, err := strconv.Atoi(os.Getenv("MAP_HEIGHT"))

	if err != nil {
		panic("failed to parse MAP_HEIGHT")
	}

	parsedMapWidth, err := strconv.Atoi(os.Getenv("MAP_WIDTH"))

	if err != nil {
		panic("failed to parse MAP_WIDTH")
	}

	config = &Config{
		DefaultMaxPlayers:       parsedDefaultMaxPlayers,
		HashCost:                parsedHashCost,
		JwtDurationInMinutes:    parsedJwtDurationInMinutes,
		JwtRenewDueInMinutes:    parsedJwtRenewDueInMinutes,
		JwtSecret:               os.Getenv("JWT_SECRET"),
		MapHeight:               parsedMapHeight,
		MapWidth:                parsedMapWidth,
		ServerAddress:           os.Getenv("SERVER_ADDRESS"),
		SQLiteDatabasePath:      os.Getenv("SQLITE_DATABASE_PATH"),
		WebSocketAllowedOrigins: os.Getenv("WEBSOCKET_ALLOWED_ORIGINS"),
	}
}
