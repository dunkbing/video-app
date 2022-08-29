package configs

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	DbUri string
}

var config *Config

func LoadEnv() {
	_ = godotenv.Load()
	if config == nil {
		config = &Config{
			DbUri: os.Getenv("MONGO_URI"),
		}
	}
}

func GetConfig() *Config {
	return config
}
