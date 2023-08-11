package config

import "time"

type Config struct {
	Database Database
	Redis    Redis
	API      API
}

type API struct {
	Address string `env:"API_ADDRESS" default:":9000"`
}
type Redis struct {
	Address  string        `env:"REDIS_ADDRESS" default:":6379"`
	Password string        `env:"REDIS_PASSWORD"`
	Expire   time.Duration `env:"REDIS_EXPIRE" default:"1h"`
}
type Database struct {
	DSN string `env:"DATA_SOURCE_NAME" default:"user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True"`
}
