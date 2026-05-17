package common

type Config struct {
	DbUrl string
	Port  string
}

func NewConfig() *Config {
	return &Config{
		DbUrl: "localhost:5432",
		Port:  "8000",
	}
}
