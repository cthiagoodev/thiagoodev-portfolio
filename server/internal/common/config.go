package common

import "os"

type Config struct {
	DbUrl string
	Port  string
}

func NewConfig() *Config {
	dbUrl := os.Getenv("DATABASE_URL")

	return &Config{
		DbUrl: dbUrl,
		Port:  "8000",
	}
}
