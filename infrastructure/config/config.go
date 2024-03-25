package config

type Config struct {
	SQLiteDatabasePath string
}

var config = Config{
	SQLiteDatabasePath: "database.sqlite3",
}

func GetConfig() Config {
	return config
}
