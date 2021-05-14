package config

var isDBConnected = false
var isMongoConnected = false
var isRedisConnected = false

type Config struct {
	Enabled      bool
	DatabasePath string
	Port         string
}

func NewConfig() *Config {
	return &Config{
		Enabled:      true,
		DatabasePath: "./example.db",
		Port:         "8000",
	}
}

// func ConnectDB(config *Config) (*sql.DB, error)

// func ConnectMongo(config *Config) (*sql.DB, error) {
// 	return sql.Open("sqlite3", config.DatabasePath)
// }

// func ConnectRedis(config *Config) (*sql.DB, error) {
// 	return sql.Open("sqlite3", config.DatabasePath)
// }
