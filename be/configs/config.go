package configs

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	DbUri      string
	CorsOrigin string
}

var config *Config

func LoadEnv() {
	_ = godotenv.Load()
	if config == nil {
		config = &Config{
			DbUri:      os.Getenv("MONGO_URI"),
			CorsOrigin: os.Getenv("CORS_ORIGIN"),
		}
	}
}

func GetConfig() *Config {
	return config
}
