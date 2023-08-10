package config

type Config struct {
	Database Database
	API      API
}

type API struct {
	Address string `env:"API_ADDRESS" default:":9000"`
}

type Database struct {
	DSN string `env:"DATA_SOURCE_NAME" default:"user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True"`
}
