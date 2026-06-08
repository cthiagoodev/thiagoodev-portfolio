package common

import "os"

type Config struct {
	DbUrl string
	Port  string
}

func NewConfig() *Config {
	dbUrl := os.Getenv("DATABASE_URL")
	port := os.Getenv("PORT")

	return &Config{
		DbUrl: dbUrl,
		Port:  port,
	}
}
